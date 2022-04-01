package schedule

import (
	"aperia/config"
	"aperia/util"
	"aperia/util/log"
	"github.com/robfig/cron/v3"
)

func moeTask() {
	c := cron.New()
	// 吃饭
	msg1 := "[CQ:image,file=https://pic.imgdb.cn/item/5fd752c13ffa7d37b3b49a45.jpg]"
	_, err := c.AddFunc("0 12,17 * * *", func() {
		util.HttpGet(host + "/send_group_msg?group_id=" + config.GroupID + "&message=" + msg1)
	})
	if err != nil {
		log.Logger.Error("cron task1 err")
	}
	// 学习
	msg2 := "[CQ:image,file=https://pic.imgdb.cn/item/5fd752c13ffa7d37b3b49a45.jpg]"
	_, err = c.AddFunc("0 8 * * *", func() {
		util.HttpGet(host + "/send_group_msg?group_id=" + config.GroupID + "&message=" + msg2)
	})
	if err != nil {
		log.Logger.Error("cron task2 err")
	}
	// 刷牙
	msg3 := "[CQ:image,file=https://pic.imgdb.cn/item/5fd752c13ffa7d37b3b49a45.jpg]"
	_, err = c.AddFunc("0 23 * * *", func() {
		util.HttpGet(host + "/send_group_msg?group_id=" + config.GroupID + "&message=" + msg3)
	})
	if err != nil {
		log.Logger.Error("cron task2 err")
	}
	// 睡觉
	msg4 := "[CQ:image,file=https://pic.imgdb.cn/item/5fd752c13ffa7d37b3b49a45.jpg]"
	_, err = c.AddFunc("0 0 * * *", func() {
		util.HttpGet(host + "/send_group_msg?group_id=" + config.GroupID + "&message=" + msg4)
	})
	if err != nil {
		log.Logger.Error("cron task2 err")
	}
}
