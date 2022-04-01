package main

import (
	"aperia/config"
	"aperia/controller"
	"aperia/function/schedule"
	"fmt"
	"net/http"
	"os"
)

func main() {
	controller.StartWebsocket()
	err := http.ListenAndServe(config.Addr+`:`+config.WsPort, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	go schedule.Start()
}
