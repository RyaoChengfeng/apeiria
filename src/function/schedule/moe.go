package schedule

import (
	"apeiria/config"
	"apeiria/util"
	"apeiria/util/log"
	"github.com/robfig/cron/v3"
)

func moeTask() {
	c := cron.New()
	// 吃饭
	groups := config.C.QQ.GroupList
	msg1 := "[CQ:image,file=https://pic.imgdb.cn/item/5fd752c13ffa7d37b3b49a45.jpg]"
	_, err := c.AddFunc("0 12,17 * * *", func() {
		moeSend(groups, msg1)
	})
	if err != nil {
		log.Logger.Error("cron task1 err")
	}
	// 学习
	msg2 := "[CQ:image,file=https://pic.imgdb.cn/item/5fd752c13ffa7d37b3b49a45.jpg]"
	_, err = c.AddFunc("0 8 * * *", func() {
		moeSend(groups, msg2)
	})
	if err != nil {
		log.Logger.Error("cron task2 err")
	}
	// 刷牙
	msg3 := "[CQ:image,file=https://pic.imgdb.cn/item/5fd752c13ffa7d37b3b49a45.jpg]"
	_, err = c.AddFunc("0 23 * * *", func() {
		moeSend(groups, msg3)
	})
	if err != nil {
		log.Logger.Error("cron task2 err")
	}
	// 睡觉
	msg4 := "[CQ:image,file=https://pic.imgdb.cn/item/5fd752c13ffa7d37b3b49a45.jpg]"
	_, err = c.AddFunc("0 0 * * *", func() {
		moeSend(groups, msg4)
	})
	if err != nil {
		log.Logger.Error("cron task2 err")
	}
}

func moeSend(groups []string, msg string) {
	for _, group := range groups {
		util.SendGroup(group, msg)
	}
}
