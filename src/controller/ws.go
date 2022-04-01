package controller

import (
	"apeiria/config"
	"apeiria/util/log"
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

const (
	// 允许等待的写入时间
	WriteWait = 10 * time.Second
	// Time allowed to read the next pong message from the peer.
	PongWait = 999999 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	PingPeriod = (PongWait * 9) / 10
	// Maximum message size allowed from peer.
	MaxMessageSize = 51200
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 允许所有的CORS 跨域请求，正式环境可以关闭
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 最大的连接ID，每次连接都加1 处理
var maxConnId int64

// 客户端读写消息
type wsMessage struct {
	// websocket.TextMessage 消息类型
	messageType int
	data        []byte
}

// ws 的所有连接
// 用于广播
var WsConnAll map[int64]*wsConnection

// 客户端连接
type wsConnection struct {
	wsSocket *websocket.Conn // 底层websocket
	inChan   chan *wsMessage // 读队列
	outChan  chan *wsMessage // 写队列

	mutex     sync.Mutex // 避免重复关闭管道,加锁处理
	isClosed  bool
	closeChan chan byte // 关闭通知
	id        int64
}

func wsHandler(rsp http.ResponseWriter, req *http.Request) {
	// 应答客户端告知升级连接为websocket
	wsSocket, err := upgrader.Upgrade(rsp, req, nil)
	if err != nil {
		log.Logger.Error("升级为websocket失败", err.Error())
	}
	maxConnId++
	log.Logger.Info(wsSocket.RemoteAddr(), maxConnId)

	// 连接数保持一定数量，超过的部分不提供服务
	// 如果要控制连接数可以计算WsConnAll长度 len(WsConnAll)
	wsConn := &wsConnection{
		wsSocket:  wsSocket,
		inChan:    make(chan *wsMessage, 1000),
		outChan:   make(chan *wsMessage, 1000),
		closeChan: make(chan byte),
		isClosed:  false,
		id:        maxConnId,
	}

	WsConnAll[maxConnId] = wsConn
	// 处理器,发送定时信息，避免意外关闭
	go processLoop(wsConn)
	// 读协程panic: runtime error: invalid memory address or nil pointer dereference

	go wsReadLoop(wsConn)
	// 写协程
	go wsWriteLoop(wsConn)
}

// 读取消息队列中的消息
func (wsConn *wsConnection) wsRead() (*wsMessage, error) {
	select {
	case msg := <-wsConn.inChan:
		// 获取到消息队列中的消息
		return msg, nil
	case <-wsConn.closeChan:
		// 应该传空结构体而不是nil，用nil会导致msg.data报错
		//return nil, errors.New("连接已经关闭")
		return &wsMessage{}, errors.New("连接已经关闭")
	}
}

// 写入消息到队列中
func (wsConn *wsConnection) wsWrite(messageType int, data []byte) error {
	select {
	case wsConn.outChan <- &wsMessage{messageType, data}:
	case <-wsConn.closeChan:
		return errors.New("连接已经关闭")
	}
	return nil
}

// 处理队列中的消息
func processLoop(wsConn *wsConnection) {
	// 处理消息队列中的消息
	// 获取到消息队列中的消息，处理完成后，发送消息给客户端
	for {
		msg, err := wsConn.wsRead()
		if err != nil {
			log.Logger.Error("获取消息出现错误:", err.Error())
			wsConn.close()
			return
		}
		// log.Println(msg.messageType)
		// log.Println(string(msg.data))
		var msgData map[string]interface{}
		err = json.Unmarshal(msg.data, &msgData)
		if err != nil {
			log.Logger.Error("json信息解析错误", err.Error())
		}
		log.Logger.Debug("收到消息：", msgData)
		go HandleWsMsg(msgData)
	}
}

// 处理消息队列中的消息
func wsReadLoop(wsConn *wsConnection) {
	// 设置消息的最大长度
	wsConn.wsSocket.SetReadLimit(MaxMessageSize)
	err := wsConn.wsSocket.SetReadDeadline(time.Now().Add(PongWait))
	if err != nil {
		log.Logger.Error("wsSocket.SetReadDeadline failed", err.Error())
	}
	for {
		msgType, data, err := wsConn.wsSocket.ReadMessage()
		if err != nil {
			websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure)
			log.Logger.Error("消息读取出现错误", err.Error())
			wsConn.close()
			// TODO 换成break和用return的区别？
			return
			//break
		}
		req := &wsMessage{
			msgType,
			data,
		}
		// 放入请求队列,消息入栈
		select {
		case wsConn.inChan <- req:
		case <-wsConn.closeChan:
			return
			//break
		}
	}
}

// 发送消息给客户端
func wsWriteLoop(wsConn *wsConnection) {
	ticker := time.NewTicker(PingPeriod)
	defer func() {
		ticker.Stop()
	}()
	for {
		select {
		// 取一个应答
		case msg := <-wsConn.outChan:
			// 写给websocket
			if err := wsConn.wsSocket.WriteMessage(msg.messageType, msg.data); err != nil {
				log.Logger.Error("发送消息给客户端发生错误", err.Error())
				// 切断服务
				wsConn.close()
				return
				//break
			}
		case <-wsConn.closeChan:
			// 获取到关闭通知
			return
			//break
		case <-ticker.C:
			// 出现超时情况
			err := wsConn.wsSocket.SetWriteDeadline(time.Now().Add(WriteWait))
			if err != nil {
				return
			}
			err = wsConn.wsSocket.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				return
			}
		}
	}
}

// 关闭连接
func (wsConn *wsConnection) close() {
	log.Logger.Debug("调用ws.close()")
	wsConn.wsSocket.Close()
	wsConn.mutex.Lock()
	defer wsConn.mutex.Unlock()
	if wsConn.isClosed == false {
		wsConn.isClosed = true
		// 删除这个连接的变量
		delete(WsConnAll, wsConn.id)
		close(wsConn.closeChan)
	}
}

// 启动程序
func StartWebsocket() {
	WsConnAll = make(map[int64]*wsConnection)
	http.HandleFunc("/", wsHandler)
	err := http.ListenAndServe(config.C.Bot.Addr+`:`+config.C.Bot.WsPort, nil)
	if err != nil {
		log.Logger.Error(err)
		// 重启服务
		//StartWebsocket()
	}
}
