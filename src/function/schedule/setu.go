package schedule

import (
	"aperia/config"
	"aperia/util"
	"aperia/util/log"
	"github.com/robfig/cron/v3"
)

func mainSetu() {
	c := cron.New()
	_, err := c.AddFunc("0 2 * * *", func() {
		url := util.GetSetu()
		msg := "深夜涩图[CQ:image,file=" + url + "]"
		util.HttpGet(host + "/send_group_msg?group_id=" + config.GroupID + "&message=" + msg)
	})
	if err != nil {
		log.Logger.Error("setu err")
	}

	_, err = c.AddFunc("0 21 * * [1-5]", func() {
		url := util.GetSetu()
		msg := "今天学习辛苦了！来张涩图提提神吧～～！[CQ:image,file=" + url + "]"
		util.HttpGet(host + "/send_group_msg?group_id=" + config.GroupID + "&message=" + msg)
	})
	if err != nil {
		log.Logger.Error("setu err")
	}
}
