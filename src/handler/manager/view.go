package manager

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Manager(c *gin.Context) {
	v, ok := c.Get("username")
	if !ok {

	}
	username := v.(string)

	if username != "localhost" {
		c.String(http.StatusOK, username+"禁止進入")
		return
	}

	c.HTML(http.StatusOK, "manager.html", gin.H{
		"title": "後台",
	})
}
