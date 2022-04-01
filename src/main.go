package main

import (
	"apeiria/config"
	"apeiria/controller"
	"apeiria/function/schedule"
	"fmt"
	"net/http"
	"os"
)

func main() {
	controller.StartWebsocket()
	err := http.ListenAndServe(config.C.Bot.Addr+`:`+config.C.Bot.WsPort, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	go schedule.Start()
}
