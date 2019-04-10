package handler

import (
	"orderfood/src/handler/manager"
	"orderfood/src/handler/middleware"
	"orderfood/src/handler/user"
	"orderfood/src/handler/ws"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("src/templates/*.html")
	router.StaticFile("src/img/rice.jpg", "src/img/rice.jpg")
	router.StaticFile("src/img/vag.jpg", "src/img/vag.jpg")
	router.StaticFile("favicon.ico", "src/img/head.ico")

	router.StaticFile("src/css/style.css", "src/css/style.css")
	router.StaticFile("src/css/managerStyle.css", "src/css/managerStyle.css")

	router.StaticFile("src/js/menu.js", "src/js/menu.js")
	router.StaticFile("src/js/menuView.js", "src/js/menuView.js")
	router.StaticFile("src/js/post.js", "src/js/post.js")
	router.StaticFile("src/js/manager.js", "src/js/manager.js")
	router.StaticFile("src/js/websocket.js", "src/js/websocket.js")

	admin := router.Group("")
	ver := admin.Use(
		middleware.Verify,
	)

	ver.GET("", user.Index)
	ver.POST("/get/menu", user.GetMenu)
	ver.POST("/post/order", user.Order)
	ver.POST("/get/user/orders", user.UserOrder)
	ver.GET("/manager", manager.Manager)
	ver.POST("/get/name", user.GetUserName)
	ver.PUT("/user", user.ModifyUser)

	router.POST("/get/order", manager.GetTotalOrders)
	router.POST("/post/view", manager.ChangeView)

	router.GET("/ws", ws.Connect)

	return router
}
