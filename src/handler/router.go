package handler

import (
	"orderfood/src/handler/auth"
	"orderfood/src/handler/manager"
	managerMenu "orderfood/src/handler/manager/menu"
	managerView "orderfood/src/handler/manager/view"
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
	router.StaticFile("src/js/models/reqs", "src/js/models/reqs")
	router.StaticFile("src/js/menu.js", "src/js/menu.js")
	router.StaticFile("src/js/menuView.js", "src/js/menuView.js")
	router.StaticFile("src/js/post.js", "src/js/post.js")
	router.StaticFile("src/js/ajax.js", "src/js/ajax.js")
	router.StaticFile("src/js/manager/treenode.js", "src/js/manager/treenode.js")
	router.StaticFile("src/js/manager/api.js", "src/js/manager/api.js")
	router.StaticFile("src/js/manager/home/main.js", "src/js/manager/home/main.js")
	router.StaticFile("src/js/manager/manageMenu/main.js", "src/js/manager/manageMenu/main.js")
	router.StaticFile("src/js/manager/newshop/main.js", "src/js/manager/newshop/main.js")
	router.StaticFile("src/js/manager/newoption/main.js", "src/js/manager/newoption/main.js")

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
	mangr.GET("/newoption", managerView.NewOption)

	api := router.Group("api")

	// manager api
	apiManager := api.Group("/manager")
	apiManager.Use(
		middleware.Verify,
	)

	apiManager.PUT("/changeshop", manager.ChangeView)

	managerMenuGroup := apiManager.Group("/menu")
	// manager menu
	managerMenuGroup.GET("", user.GetMenu)
	managerMenuGroup.GET("/menu/:shop", managerMenu.GetMenu)

	managerMenuGroup.GET("/shopmenu/:shopID", managerMenu.GetShopMenu)

	managerMenuGroup.POST("/shop", managerMenu.AddShop)
	managerMenuGroup.GET("/shop", managerMenu.GetShop)
	managerMenuGroup.PUT("/shop/:id", managerMenu.UpdateShop)
	managerMenuGroup.DELETE("/shop/:id", managerMenu.DeleteShop)

	managerMenuGroup.POST("/item", managerMenu.AddItem)
	managerMenuGroup.GET("/item/:shopID", managerMenu.GetItem)
	managerMenuGroup.PUT("/item/:id", managerMenu.UpdateItem)
	managerMenuGroup.DELETE("/item/:id", managerMenu.DeleteItem)

	managerMenuGroup.POST("/itemoption", managerMenu.AddItemOption)
	managerMenuGroup.DELETE("/itemoption/:id", managerMenu.DeleteItemOption)

	managerMenuGroup.POST("/option", manager.AddOption)
	managerMenuGroup.PUT("/option/:id", managerMenu.UpdateOption)
	managerMenuGroup.DELETE("/option/:id", managerMenu.DeleteOption)

	managerMenuGroup.POST("/selection", managerMenu.AddSelection)
	managerMenuGroup.PUT("/selection/:id", managerMenu.UpdateSelection)
	managerMenuGroup.DELETE("/selection/:id", managerMenu.DeleteSelection)

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
