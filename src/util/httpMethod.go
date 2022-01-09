package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"aperia/util/log"
)

func HttpGet(url string) *http.Response {
	defer func() {
		info := recover()
		if info != nil {
			log.Logger.Info("recover from http.Get",info)
		}
	}()
	rsp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()
	return rsp
}

func HttpPost(url string,data interface{}) *http.Response {
	defer func() {
		info := recover()
		if info != nil {
			log.Logger.Info("recover from http.Post",info)
		}
	}()
	jsonStr, _ := json.Marshal(data)
	rsp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()
	return rsp
}

func GetRspBody(rsp *http.Response) []byte {
	body, _ := ioutil.ReadAll(rsp.Body)
	return body
}