package manager

import (
	"net/http"
	"orderfood/src/handler/auth"
	"orderfood/src/logic"

	"github.com/gin-gonic/gin"
)

func Manager(c *gin.Context) {
	user := auth.MapUserName(c)
	if user != "哥" {
		c.String(http.StatusOK, user+"禁止進入")
		return
	}

	c.HTML(http.StatusOK, "manager.html", gin.H{
		"title": "後台",
	})
}

func ChangeView(c *gin.Context) {
	view := c.PostForm("view")

	logic.SetView(view)

	c.String(http.StatusOK, view)
}

func GetTotalOrders(c *gin.Context) {
	totalList := logic.IntegrationOrders()

	c.String(http.StatusOK, totalList)
}
