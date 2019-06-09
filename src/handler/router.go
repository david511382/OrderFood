package handler

import (
	"orderfood/src/handler/auth"
	"orderfood/src/handler/manager"
	"orderfood/src/handler/middleware"
	"orderfood/src/handler/order"
	"orderfood/src/handler/menu"
	"orderfood/src/handler/swag"
	"orderfood/src/handler/user"
	"orderfood/src/handler/ws"

	"github.com/gin-gonic/gin"
)

func Init(isReleaseMode bool) *gin.Engine {
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

	// manager
	mangr := router.Group("manager")
	mangr.Use(
		middleware.Verify,
	)
	mangr.GET("/", manager.Manager)
	mangr.PUT("/shop", manager.ChangeView)

	api := router.Group("api")

	// menu
	me := api.Group("/menu")
	me.Use(
		middleware.Verify,
	)

	me.GET("", user.GetMenu)
	me.GET("/menu/:shop", menu.GetMenu)

	me.POST("/shop", menu.AddShop)
	me.GET("/shop", menu.GetShop)
	me.PUT("/shop", menu.UpdateShop)
	me.DELETE("/shop", menu.DeleteShop)

	me.POST("/item", menu.AddItem)
	me.GET("/item/:shop_id", menu.GetItem)
	me.PUT("/item", menu.UpdateItem)
	me.DELETE("/item", menu.DeleteItem)

	me.POST("/itemOption", menu.AddItemOption)
	me.GET("/itemOption", menu.GetItemOption)
	me.PUT("/itemOption", menu.UpdateItemOption)
	me.DELETE("/itemOption", menu.DeleteItemOption)

	me.POST("/option", menu.AddOption)
	me.GET("/option", menu.GetOption)
	me.PUT("/option", menu.UpdateOption)
	me.DELETE("/option", menu.DeleteOption)

	me.POST("/optionSelection", menu.AddOptionSelection)
	me.GET("/optionSelection", menu.GetOptionSelection)
	me.PUT("/optionSelection", menu.UpdateOptionSelection)
	me.DELETE("/optionSelection", menu.DeleteOptionSelection)

	me.POST("/selection", menu.AddSelection)
	me.GET("/selection", menu.GetSelection)
	me.PUT("/selection", menu.UpdateSelection)
	me.DELETE("/selection", menu.DeleteSelection)

	// auth
	au := api.Group("auth")
	au.POST("/register", auth.Register)

	// user
	usr := api.Group("user")
	usr.Use(
		middleware.Verify,
	)
	usr.GET("/", user.GetUserName)
	usr.PUT("/", user.ModifyUser)

	// order
	odr := api.Group("order")
	odr.Use(
		middleware.Verify,
	)
	odr.GET("/", order.UserOrder)
	odr.GET("/all", order.GetTotalOrders)
	odr.PUT("/", order.Order)

	return router
}
