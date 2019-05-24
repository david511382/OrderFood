//+build !release,!docker

package tags

import (
	_ "orderfood/docs"

	"github.com/gin-gonic/gin"
)

func initConfig()  {
	isReleaseMode = false
	configPath = "./src/config/config.yml"
}

func run(router *gin.Engine, addr string) {
	if err := router.Run(addr); err != nil {
		panic(err)
	}
}
