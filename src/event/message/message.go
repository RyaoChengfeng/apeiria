package message

import (
	"apeiria/config"
	"apeiria/function/setu"
	util2 "apeiria/util"
	"apeiria/util/log"
	"strconv"
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
	userID := strconv.FormatFloat(message["user_id"].(float64), 'f', -1, 64)
	if util2.IsInStringList(userID, config.C.QQ.SuperUserList) {
		return
	}
	funcName, _ := CheckPrivateMessage(message["raw_message"].(string))
	switch funcName {
	case "about":
		util2.SendPrivate(userID, "hello")
	//case `yunbanke`:
	//	util2.SendPrivate(message["user_id"].(float64), "请输入您的账号和密码以开始签到, 格式为：\\云班课自动签到 账号:密码:持续时间")
	//case `start_yunbanke`:
	//	util2.SendPrivate(message["user_id"].(float64), "开始自动签到")
	////yunbanke.YunBankeCheckIn(data[0],data[1],data[2])
	case `nil`:
		util2.SendPrivate(userID, config.C.Bot.Name+"目前不清楚你在说什么哦")
	}
}

func handleGroup(message map[string]interface{}) {
	log.Logger.Debug("群组消息：", message["raw_message"])
	funcName, _ := CheckGroupMessage(message["raw_message"].(string))
	groupID := strconv.FormatFloat(message["group_id"].(float64), 'f', -1, 64)
	//fmt.Println("GroupID:", groupID)
	if util2.IsInStringList(groupID, config.C.QQ.GroupList) {
		//fmt.Println("funcName:", funcName)
		switch funcName {
		case "setu":
			log.Logger.Debug("发送图片到：", groupID)
			setu.SendGroupSetu(groupID)
		case "echo":
			util2.SendGroup(groupID, message["raw_message"].(string))
		}
	}
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
		if util2.CheckWordExist(`\云班课自动签到 `, msgStr) {
			if util2.CheckRegexpMatch(`\云班课自动签到 (.*):(.*):(.*)`, msgStr) {
				data := util2.FindMatch(`\云班课自动签到 (.*):(.*):(.*)`, msgStr)
				return "start_yunbanke", data
			}
		}
		return "nil", nil
	}
}

// CheckPrivatemessage 检查群组信息，返回对应的功能名称
func CheckGroupMessage(msgStr string) (string, []string) {
	switch msgStr {
	case `\help`, `\帮助`:
		return "help", nil
	case `\作者`, `\author`:
		return "author", nil
	case `\info`, `\information`, `\about`:
		return "about", nil
	default:
		if util2.CheckRegexpMatch("来[张点][色涩]图|[涩色]图来|想要[涩色]图|[涩色]图[Tt][Ii][Mm][Ee]", msgStr) {
			return "setu", nil
		} else if util2.CheckRegexpMatch("echo", msgStr) {
			return "echo", nil
		}
		return "nil", nil
	}
}
