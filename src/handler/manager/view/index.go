package view

import (
	"net/http"
	"orderfood/src/logic"
	"strconv"

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

func MenuTree(c *gin.Context) {
	view, err := logic.MenuTreeView()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, view)
}

func NewShop(c *gin.Context) {
	view, err := logic.NewShopView()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, view)
}

func NewOption(c *gin.Context) {
	view, err := logic.NewOptionView()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, view)
}

func ManageMenu(c *gin.Context) {
	// v, ok := c.Get("username")
	// if !ok {
	//     c.AbortWithError(http.StatusBadRequest, nil)
	// 	return
	// }
	// username := v.(string)

	shopIDStr := c.Query("shopID")
	shopID, err := strconv.Atoi(shopIDStr)
	if err != nil {
		shopID = 0
	}

	view, err := logic.ManageMenuView(int32(shopID))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// c.HTML(http.StatusOK, "manager.html", gin.H{
	// 	"title": "後台",
	// })

	c.JSON(http.StatusOK, view)
	// c.Writer.Write([]byte(view))
	// c.Writer.WriteHeader(http.StatusOK)
}
