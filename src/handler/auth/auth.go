package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserName(c *gin.Context) {
	v, ok := c.Get("username")
	if !ok{

	}
	name := v.(string)

	c.String(http.StatusOK, name)
}