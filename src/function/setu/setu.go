package setu

import (
	"aperia/config"
	"aperia/util"
	"fmt"
)

var host = fmt.Sprintf(config.Addr + `:` + config.BotPort)

func SendGroupSetu(groupID string) {
	url := util.GetSetu()
	msg := "[CQ:image,file=" + url + "]"
	util.HttpGet(host + "/send_group_msg?group_id=" + groupID + "&message=" + msg)
}
