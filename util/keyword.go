package util

import (
	"regexp"
	"strings"
)

// CheckPrivatemessage 检查私聊信息，返回对应的功能名称
func CheckPrivatemessage(msgStr string) (string,[]string) {
	switch msgStr {
	case `\help`, `\帮助`:
		return "help",nil
	case `\云班课自动签到`, `\云班课`:
		return "yunbanke",nil
	case `\作者`,`\author`:
		return "author",nil
	case `\info`,`\information`,`\about`:
		return "about",nil
	default:
		if CheckWordExist(`\云班课自动签到 `,msgStr) {
			if CheckMatch(`\云班课自动签到 (.*):(.*):(.*)`,msgStr) {
				data := FindMatch(`\云班课自动签到 (.*):(.*):(.*)`,msgStr)
				return "start_yunbanke",data
			}
		}
		return "nil",nil
	}
}

// CheckPrivatemessage 检查群组信息，返回对应的功能名称
func CheckGroupmessage(msgStr string) string  {
	switch msgStr {
	case `\help`, `\帮助`:
		return "help"
	case `\作者`,`\author`:
		return "author"
	case `\info`,`\information`,`\about`:
		return "about"
	default:
		return "nil"
	}
}

func FindMatch(key string,msgStr string) []string {
	reg := regexp.MustCompile(key)
	return reg.FindAllString(msgStr, -1)
}

func CheckMatch(key string,msgStr string) bool {
	match, _ := regexp.MatchString(key, msgStr)
	return match
}

func CheckWordExist(key string,msgStr string) bool {
	return strings.Contains(msgStr,key)
}
