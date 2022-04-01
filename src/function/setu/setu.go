package setu

import (
	"apeiria/util"
	"apeiria/util/log"
)

func SendGroupSetu(groupID string) {
	url := util.GetSetu()
	path, err := util.DownloadImage(url)
	log.Logger.Debug(url)
	if err != nil {
		log.Logger.Warn(err)
		return
	}

	log.Logger.Debug("Sending setu to group %s,url:%s", groupID, url)
	if url == "" {
		util.SendGroup(groupID, "高性能机器人没有偷到涩图哦~")
		return
	}
	msg := "[CQ:image,file=file://" + path + ",url=" + "file://" + path + "]"
	//msg := "[CQ:image,file=" + url + ",url=" + url + "]"
	util.SendGroup(groupID, msg)
	err = util.DeleteFile(path)
	log.Logger.Debug("Delete file:", path)
	if err != nil {
		log.Logger.Warn("Delete file failed:", path)
	}
}
