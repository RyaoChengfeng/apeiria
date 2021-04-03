package message

import (
	"rinqqbot/util"
	"rinqqbot/util/log"
)

func CheckType(msg map[string]interface{}) {
	switch msg["message_type"] {
	case "private":
		log.Logger.Debug(msg)
		handlePrivate(msg)
		break
	case "group":
		handleGroup(msg)
		break
	default:
		break
	}
}

func handlePrivate(message map[string]interface{}) {
	log.Logger.Debug("私聊消息",message["raw_message"])
	funcName,_ := util.CheckPrivatemessage(message["raw_message"].(string))
	switch funcName {
	case "about":
		util.SendPrivate(message["user_id"].(float64),"hello")
	case `yunbanke`:
		util.SendPrivate(message["user_id"].(float64),"请输入您的账号和密码以开始签到, 格式为：\\云班课自动签到 账号:密码:持续时间")
	case `start_yunbanke`:
		util.SendPrivate(message["user_id"].(float64),"开始自动签到")
		//yunbanke.YunBankeCheckIn(data[0],data[1],data[2])
	}
}

func handleGroup(message map[string]interface{}) {
	log.Logger.Debug("群组消息",message["raw_message"])
}