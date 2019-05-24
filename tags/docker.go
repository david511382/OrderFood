//+build docker

package tags

import (
	"github.com/gin-gonic/gin"
)

func initConfig() {
	isReleaseMode = true
	configPath = "./src/config/docker-config.yml"
}

func run(router *gin.Engine, addr string) {
	if err := router.Run(addr); err != nil {
		panic(err)
	}
}
