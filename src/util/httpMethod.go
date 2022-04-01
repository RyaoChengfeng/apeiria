package util

import (
	"aperia/util/log"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func HttpGet(url string) *http.Response {
	defer func() {
		info := recover()
		if info != nil {
			log.Logger.Info("recover from http.Get", info)
		}
	}()
	rsp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	//defer rsp.Body.Close()
	return rsp
}

func HttpPost(url string, data interface{}) *http.Response {
	defer func() {
		info := recover()
		if info != nil {
			log.Logger.Info("recover from http.Post", info)
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

func HttpProxyGet(URL string, proxy string) *http.Response {
	uri, err := url.Parse(proxy)
	if err != nil {
		log.Logger.Fatal("parse url error: ", err)
	}

	client := http.Client{
		Transport: &http.Transport{
			// 设置代理
			Proxy: http.ProxyURL(uri),
		},
	}
	defer func() {
		info := recover()
		if info != nil {
			log.Logger.Info("recover from http.Get", info)
		}
	}()
	rsp, err := client.Get(URL)
	if err != nil {
		panic(err)
	}
	//defer rsp.Body.Close()
	return rsp
}

func GetRspBody(rsp *http.Response) []byte {
	body, _ := ioutil.ReadAll(rsp.Body)
	return body
}
