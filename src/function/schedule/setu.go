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
		msg := "深夜涩图"
		setu(config.C.QQ.GroupList, msg)
	})
	if err != nil {
		log.Logger.Error("setu err")
	}

	_, err = c.AddFunc("0 21 * * [1-5]", func() {
		msg := "今天学习辛苦了！来张涩图提提神吧～～！"
		setu(config.C.QQ.GroupList, msg)
	})
	if err != nil {
		log.Logger.Error("setu err")
	}
}

func setu(groups []string, cmt string) {
	url := util.GetSetu()
	path, err := util.DownloadImage(url)
	log.Logger.Debug(url)
	if err != nil {
		log.Logger.Warn(err)
		return
	}

	log.Logger.Debug("Sending setu to groups: %s, url:%s", groups, url)
	if url == "" {
		log.Logger.Warn("setu url is empty")
		return
	}
	msg := cmt + "[CQ:image,file=file://" + path + ",url=" + "file://" + path + "]"
	for _, group := range groups {
		util.SendGroup(group, msg)
	}
	err = util.DeleteFile(path)
	log.Logger.Debug("Delete file:", path)
	if err != nil {
		log.Logger.Warn("Delete file failed:", path)
	}
}
