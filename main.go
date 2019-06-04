package main

import (
	"orderfood/src/handler"
	"orderfood/src/logic"
	"orderfood/tags"
)

// @title Order Food API
// @version 1.0
// @description 訂餐系統

// @BasePath /api/

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfg, isReleaseMode := tags.InitConfig("")

	logic.Init(cfg)

	router := handler.Init(isReleaseMode)

	addr := cfg.Domain()

	tags.Run(router, addr)
}
