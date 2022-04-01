package util

import (
	"apeiria/config"
	"apeiria/util/log"
	"encoding/json"
	"os"
	"path"
	"path/filepath"
)

func GetSetu() string {
	r := make(map[string]interface{})
	jsonErr := json.Unmarshal(GetRspBody(HttpGet(config.C.Setu.API+"?r18="+config.C.Setu.R18)), &r)
	if jsonErr != nil {
		log.Logger.Error(jsonErr)
	}
	log.Logger.Debug(r)
	if len(r["data"].([]interface{})) == 0 {
		return ""
	}
	url := r["data"].([]interface{})[0].(map[string]interface{})["urls"].(map[string]interface{})["original"].(string)
	return url
}

func DownloadImage(imgUrl string) (filename string, err error) {
	appPath, _ := os.Getwd()
	fileBaseName := path.Base(imgUrl)
	// 图片保存的相对路径
	imgPath := filepath.Join(config.C.Setu.ImagePath)
	filename = filepath.Join(imgPath, fileBaseName)

	body := GetRspBody(HttpProxyGet(imgUrl, config.C.Proxy))
	// 自动创建文件夹
	if err = CheckDir(appPath + imgPath); err != nil {
		return
	}
	Fpath := appPath + filename
	f, err := os.OpenFile(Fpath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		return Fpath, err
	} else {
		_, err = f.Write(body)
		if err != nil {
			return Fpath, err
		}
	}
	return Fpath, err
}

func CheckDir(path string) error {
	if _, err := os.Stat(path); err == nil {
		return nil
	} else {
		err := os.MkdirAll(path, 0711)
		if err != nil {
			return err
		}
	}
	// check again
	_, err := os.Stat(path)
	return err
}

func DeleteFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
