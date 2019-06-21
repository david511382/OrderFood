package handler

import (
	"orderfood/src/handler/auth"
	"orderfood/src/handler/manager"
	managerView "orderfood/src/handler/manager/view"
	"orderfood/src/handler/menu"
	"orderfood/src/handler/middleware"
	"orderfood/src/handler/order"
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

	// image
	router.StaticFile("src/img/rice.jpg", "src/img/rice.jpg")
	router.StaticFile("src/img/vag.jpg", "src/img/vag.jpg")
	router.StaticFile("favicon.ico", "src/img/head.ico")

	// css
	router.StaticFile("src/css/style.css", "src/css/style.css")
	router.StaticFile("css/manager.css", "src/css/manager.css")

	// js
	router.StaticFile("src/js/menu.js", "src/js/menu.js")
	router.StaticFile("src/js/menuView.js", "src/js/menuView.js")
	router.StaticFile("src/js/post.js", "src/js/post.js")
	router.StaticFile("src/js/manager.js", "src/js/manager.js")
	router.StaticFile("src/js/manager/treenode.js", "src/js/manager/treenode.js")
	router.StaticFile("src/js/manager/manageMenuMain.js", "src/js/manager/manageMenuMain.js")
	router.StaticFile("src/js/manager/api.js", "src/js/manager/api.js")
	router.StaticFile("src/js/manager/newshop.js", "src/js/manager/newshop.js")
	router.StaticFile("src/js/websocket.js", "src/js/websocket.js")

	router.GET("", user.Index)

	router.GET("/ws", ws.Connect)

	// manager
	mangr := router.Group("manager")
	mangr.Use(
		middleware.Verify,
	)
	mangr.GET("", managerView.Home)
	mangr.GET("/menutree", managerView.MenuTree)
	mangr.GET("/newshop", managerView.NewShop)
	mangr.GET("/managemenu", managerView.ManageMenu)
	mangr.PUT("/changeshop", manager.ChangeView)

	api := router.Group("api")

	// menu
	me := api.Group("/menu")
	me.Use(
		middleware.Verify,
	)

	me.GET("", user.GetMenu)
	me.GET("/menu/:shop", menu.GetMenu)

	me.GET("/shopmenu", menu.GetShopMenu)

	me.POST("/shop", menu.AddShop)
	me.GET("/shop", menu.GetShop)
	me.PUT("/shop/:id", menu.UpdateShop)
	me.DELETE("/shop/:id", menu.DeleteShop)

	me.POST("/item", menu.AddItem)
	me.GET("/item/:shopID", menu.GetItem)
	me.PUT("/item/:id", menu.UpdateItem)
	me.DELETE("/item/:id", menu.DeleteItem)

	me.POST("/itemoption", menu.AddItemOption)
	me.DELETE("/itemoption/:id", menu.DeleteItemOption)

	me.POST("/option", menu.AddOption)
	me.PUT("/option/:id", menu.UpdateOption)
	me.DELETE("/option/:id", menu.DeleteOption)

	me.POST("/selection", menu.AddSelection)
	me.PUT("/selection/:id", menu.UpdateSelection)
	me.DELETE("/selection/:id", menu.DeleteSelection)

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
