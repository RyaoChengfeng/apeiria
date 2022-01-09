package util

import (
	"aperia/config"
	"aperia/util/log"
	"encoding/json"
)

//
//func HandleFuncName(name string,msg) {
//	switch name {
//	case `yunbanke`:
//
//	}
//}

func GetSetu() string {
	r := make(map[string]interface{})
	jsonErr := json.Unmarshal(GetRspBody(HttpGet(config.SetuAPI+"?r18="+config.R18)), &r)
	if jsonErr != nil {
		log.Logger.Error(jsonErr)
	}
	url := r["data"].([]interface{})[0].(map[string]interface{})["url"].(string)
	return url
}
