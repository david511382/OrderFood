//+build !release

package main

import (
	_ "orderfood/docs"

	"github.com/gin-gonic/gin"
)

var isReleaseMode bool

func initServer() {
	isReleaseMode = false
}

func run(router *gin.Engine, addr string) {
	if err := router.Run(addr); err != nil {
		panic(err)
	}
}
