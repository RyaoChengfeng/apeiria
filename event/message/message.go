package message

import (
	"rinqqbot/config"
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
	log.Logger.Debug("私聊消息", message["raw_message"])
	funcName, _ := CheckPrivateMessage(message["raw_message"].(string))
	switch funcName {
	case "about":
		util.SendPrivate(message["user_id"].(float64), "hello")
	case `yunbanke`:
		util.SendPrivate(message["user_id"].(float64), "请输入您的账号和密码以开始签到, 格式为：\\云班课自动签到 账号:密码:持续时间")
	case `start_yunbanke`:
		util.SendPrivate(message["user_id"].(float64), "开始自动签到")
	//yunbanke.YunBankeCheckIn(data[0],data[1],data[2])
	case `nil`:
		util.SendPrivate(message["user_id"].(float64), config.BotName+"目前不清楚你在说什么哦")
	}
}

func handleGroup(message map[string]interface{}) {
	log.Logger.Debug("群组消息", message["raw_message"])
}

// CheckPrivateMessage 检查私聊信息，返回对应的功能名称
func CheckPrivateMessage(msgStr string) (string, []string) {
	switch msgStr {
	case `\help`, `\帮助`:
		return "help", nil
	case `\云班课自动签到`, `\云班课`:
		return "yunbanke", nil
	case `\作者`, `\author`:
		return "author", nil
	case `\info`, `\information`, `\about`:
		return "about", nil
	case `\hello`:
		return "hello", nil
	default:
		if util.CheckWordExist(`\云班课自动签到 `, msgStr) {
			if util.CheckMatch(`\云班课自动签到 (.*):(.*):(.*)`, msgStr) {
				data := util.FindMatch(`\云班课自动签到 (.*):(.*):(.*)`, msgStr)
				return "start_yunbanke", data
			}
		}
		return "nil", nil
	}
}

// CheckPrivatemessage 检查群组信息，返回对应的功能名称
func CheckGroupMessage(msgStr string) string {
	switch msgStr {
	case `\help`, `\帮助`:
		return "help"
	case `\作者`, `\author`:
		return "author"
	case `\info`, `\information`, `\about`:
		return "about"
	default:
		return "nil"
	}
}
