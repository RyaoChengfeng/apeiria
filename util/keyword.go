package util

import (
	"regexp"
	"strings"
)


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
