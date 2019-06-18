package middleware

import (
	"orderfood/src/database/models"
	"orderfood/src/logic"
	"strings"

	"github.com/gin-gonic/gin"
)

func Verify(c *gin.Context) {
	defer logic.LoadMembers()

	ip := c.Request.RemoteAddr
	ips := strings.Split(ip, ":")
	username := ips[0]
	if username == "[" {
		username = "localhost"
	}

	// register
	member := &models.Member{
		Name:     username,
		Username: username,
		Password: username,
	}
	err := logic.Register(member)
	if err != nil && err != logic.ExisitErr {
		return
	}

	c.Set("name", member.GetName())
	c.Set("username", member.GetUsername())
}
