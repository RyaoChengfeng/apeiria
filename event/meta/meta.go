package meta

import "rinqqbot/util/log"

func CheckType(msg map[string]interface{}) {
	switch msg["meta_event_type"] {
	case "heartbeat":
		heartbeat(msg)
		break
	case "lifecycle":
		lifecycle(msg)
		break
	default:
		break
	}
}

//心跳
func heartbeat(msg map[string]interface{}) {

}

//生命周期
func lifecycle(msg map[string]interface{}) {
	if msg["sub_type"] == "connect" {
		log.Logger.Debug("连接成功,self_id:", msg["self_id"].(float64))
	}
}
