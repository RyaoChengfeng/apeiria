package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

func init() {
	configFile := "config.yml"

	// 如果有设置 ENV ，则使用ENV中的环境
	if v, ok := os.LookupEnv("ENV"); ok {
		configFile = v + ".yml"
	}

	// 读取配置文件
	data, err := ioutil.ReadFile(fmt.Sprintf("../env/%s", configFile))

	if err != nil {
		//Logger.Println("Read config error!")
		//Logger.Panic(err)
		panic(err)
		return
	}

	config := &Config{}

	err = yaml.Unmarshal(data, config)

	if err != nil {
		//Logger.Println("Unmarshal config error!")
		//Logger.Panic(err)
		log.Print("Unmarshal config error!")
		panic(err)
		return
	}

	C = config

	//Logger.Println("Config " + configFile + " loaded.")
	log.Print("Config " + configFile + " loaded.")
	if C.Debug {
		//Logger.Printf("%+v\
		log.Printf("%+v\n", C)
	}
}
