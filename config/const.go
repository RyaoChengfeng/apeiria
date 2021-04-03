package config

import "time"

const (
	Addr = "127.0.0.1"
	Port = "22334"
	LogPath = "./log"
	LogFileName = "qqbot.log"
)

var Debug = false

const (
	// 允许等待的写入时间
	WriteWait = 10 * time.Second
	// Time allowed to read the next pong message from the peer.
	PongWait = 999999 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	PingPeriod = (PongWait * 9) / 10
	// Maximum message size allowed from peer.
	MaxMessageSize = 51200
)
