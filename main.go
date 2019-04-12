package main

import (
	"flag"
	_ "orderfood/docs"
	"orderfood/src/config"
	"orderfood/src/handler"
	"orderfood/src/logic"
)

// @title Order Food API
// @version 1.0
// @description 訂餐系統

// @host 192.168.0.144:5487
// @BasePath /api/

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
