package schedule

import (
	"aperia/config"
	"aperia/util/log"
	"fmt"
)

var host = fmt.Sprintf(config.Addr + `:` + config.BotPort)

func Start() {
	log.Logger.Info("schedule start")
	//moeTask()
	//mainSetu()
}
