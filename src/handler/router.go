package handler

import (
	"net/http"
	"orderfood/src/handler/auth"
	"orderfood/src/handler/manager"
	"orderfood/src/handler/middleware"
	"orderfood/src/handler/order"
	"orderfood/src/handler/shop"
	"orderfood/src/handler/swag"
	"orderfood/src/handler/user"
	"orderfood/src/handler/ws"

	"github.com/gin-gonic/gin"
)

func Init(isReleaseMode bool) *http.Server {
	router := gin.Default()

	if isReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
		router.GET("/docs/*any", swag.Documents)
	}

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

	router.GET("", user.Index)

	router.GET("/ws", ws.Connect)

	mangr := router.Group("manager")
	mangr.Use(
		middleware.Verify,
	)
	mangr.GET("/", manager.Manager)

	api := router.Group("api")

	api.GET("/menu", user.GetMenu)

	au := api.Group("auth")
	au.POST("/register", auth.Register)

	usr := api.Group("user")
	usr.Use(
		middleware.Verify,
	)
	usr.GET("/", user.GetUserName)
	usr.PUT("/", user.ModifyUser)

	odr := api.Group("order")
	odr.Use(
		middleware.Verify,
	)
	odr.GET("/", order.UserOrder)
	odr.GET("/all", order.GetTotalOrders)
	odr.PUT("/", order.Order)

	sop := api.Group("shop")
	sop.Use(
		middleware.Verify,
	)
	sop.PUT("/", manager.ChangeView)
	sop.POST("/", shop.AddShop)
	sop.GET("/", shop.GetShop)

	sop.POST("/item/", shop.AddItem)
	sop.GET("/item/", shop.GetItem)
	sop.POST("/size", shop.AddSize)
	sop.GET("/size", shop.GetSize)
	sop.GET("/menu", shop.GetMenu)

	server := &http.Server{
		Handler: router,
	}

	return server
}
