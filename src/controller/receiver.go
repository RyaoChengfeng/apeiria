package controller

import (
	"apeiria/event/message"
	"apeiria/event/meta"
	"apeiria/util/log"
)

func HandleWsMsg(msg map[string]interface{}) {
	// fmt.Println(msg)
	log.Logger.Debug("go协程创建")
	switch msg["post_type"] {
	case "message":
		//消息事件
		message.CheckType(msg)
		break
	//case "notice":
	//	//通知事件
	//	notice.CheckType(msg)
	//	break
	//case "request":
	//	//请求事件
	//	request.CheckType(msg)
	//	break
	case "meta_event":
		//元事件
		meta.CheckType(msg)
		break
	default:
		break
	}
	log.Logger.Debug("go协程结束")
}
