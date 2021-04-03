package main

import (
	"fmt"
	"net/http"
	"os"
	"rinqqbot/config"
	"rinqqbot/controller"
)

func main() {
	controller.StartWebsocket()
	err:= http.ListenAndServe(config.Addr+`:`+config.Port,nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


