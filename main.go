package main

import (
	"orderfood/src/config"
	"orderfood/src/handler"
	"orderfood/src/logic"
)

const (
	readTimeOut  = 10000
	writeTimeout = 10000
)

var (
	cfg *config.Config
)

// @title Order Food API
// @version 1.0
// @description 訂餐系統

// @BasePath /api/

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	initServer()

	logic.Init(cfg)

	server := handler.Init(isReleaseMode)

	addr := cfg.Domain()

	server.Addr = addr
	server.ReadTimeout = readTimeOut
	server.WriteTimeout = writeTimeout
	server.MaxHeaderBytes = 1 << 20

	run(server)
}

func init() {
	err := config.ReadConfig("./src/config/config.yml")
	if err != nil {
		panic(err)
	}

	cfg = config.Get()
}
