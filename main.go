package main

import (
	"flag"
	"orderfood/src/config"
	"orderfood/src/handler"
	"orderfood/src/logic"
)

func main() {
	flagParse()

	cfg := config.Get()

	logic.Init(cfg)

	router := handler.Init()

	router.Run(cfg.Domain())
}

func flagParse() {
	configFileName := "./config/config-develop.yml"
	flag.StringVar(&configFileName, "configfile", "./src/config/config.yml", "config path")

	// // 设置日志存储位置
	// flag.StringVar(&flagLogPath, "L", "./logs", "the dir path of path")
	// flag.StringVar(&flagLogPath, "log-path", "./logs", "the dir path of path")

	// // set config
	// flag.StringVar(&flagConfigFile, "c", "./config/config.yml", "the file path of config")
	// flag.StringVar(&flagConfigFile, "config-file", "./config/config.yml", "the file path of config")

	flag.Parse()

	err := config.ReadConfig(configFileName)
	if err != nil {
		panic(err)
	}
}
