package handler

import (
	"orderfood/src/handler/auth"
	"orderfood/src/handler/manager"
	"orderfood/src/handler/middleware"
	"orderfood/src/handler/order"
	"orderfood/src/handler/swag"
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

	router.GET("/docs/*any", swag.Documents)

	router.GET("", user.Index)

	router.GET("/ws", ws.Connect)

	mangr := router.Group("manager")
	mangrVer := mangr.Use(
		middleware.Verify,
	)
	mangrVer.GET("/", manager.Manager)

	api := router.Group("api")

	api.GET("/menu", user.GetMenu)

	au := api.Group("auth")
	au.POST("/register", auth.Register)

	usr := api.Group("user")
	usrVer := usr.Use(
		middleware.Verify,
	)
	usrVer.GET("/", user.GetUserName)
	usrVer.PUT("/", user.ModifyUser)

	odr := api.Group("order")
	odrVer := odr.Use(
		middleware.Verify,
	)
	odr.GET("/", order.UserOrder)
	odr.GET("/all", order.GetTotalOrders)
	odrVer.PUT("/", order.Order)

	sop := api.Group("shop")
	sopVer := sop.Use(
		middleware.Verify,
	)
	sopVer.PUT("/", manager.ChangeView)

	return router
}
