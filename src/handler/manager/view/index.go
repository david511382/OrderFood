package view

import (
	"net/http"
	"orderfood/src/logic"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	v, ok := c.Get("username")
	if !ok {
        c.AbortWithError(http.StatusBadRequest, nil)
		return
	}
	username := v.(string)

	data, err := logic.ManagerView(username)
	if err != nil {
		c.String(http.StatusOK, data)
		return
	}

	// c.HTML(http.StatusOK, "manager.html", gin.H{
	// 	"title": "後台",
	// })

	c.Writer.Write([]byte(data))
	c.Writer.WriteHeader(http.StatusOK)
}

func ManageShop(c *gin.Context) {
	// v, ok := c.Get("username")
	// if !ok {
    //     c.AbortWithError(http.StatusBadRequest, nil)
	// 	return
	// }
	// username := v.(string)

	shopName := c.Query("shopName")

	data, err := logic.ManageShopView(shopName)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// c.HTML(http.StatusOK, "manager.html", gin.H{
	// 	"title": "後台",
	// })

	c.Writer.Write([]byte(data))
	c.Writer.WriteHeader(http.StatusOK)
}