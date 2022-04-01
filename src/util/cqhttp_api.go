package util

import (
	"aperia/config"
	"aperia/util/log"
	"fmt"
)

var host = fmt.Sprintf(config.C.Bot.Addr + `:` + config.C.Bot.BotPort)

func SendPrivate(qq string, msg string) {
	_ = HttpGet(host + "/send_private_msg?user_id=" + qq + "&message=" + msg)
}

func SendGroup(group string, msg string) {
	log.Logger.Info("发送消息到群" + group + ":" + msg)
	_ = HttpGet(host + "/send_group_msg?group_id=" + group + "&message=" + msg)
}

func SendGroupPost(group string, msg string) {
	log.Logger.Info("发送消息到群" + group + ":" + msg)

	// var data map[string]interface{}
	data := make(map[string]interface{})
	data["group_id"] = group
	data["message"] = msg

	_ = HttpPost(host+"/send_group_msg", data)
}

func SendGift(group string, qq string, num int) {
	gift := fmt.Sprintf("[CQ:gift,qq=%s,id=%d]", qq, num)
	SendGroupPost(group, gift)
}

func SendPoke(group string, qq string) {
	poke := fmt.Sprintf("[CQ:poke,qq=%s]", qq)
	SendGroupPost(group, poke)
}
