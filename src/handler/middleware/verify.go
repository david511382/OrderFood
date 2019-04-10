package middleware

import (
	"orderfood/src/database"
	"orderfood/src/database/models"
	"orderfood/src/logic"
	"strings"

	"github.com/gin-gonic/gin"
)

func Verify(c *gin.Context) {
	defer c.Next()

	if logic.IsRquiredPassword {
		return
	}

	defer logic.LoadMembers()

	ip := c.Request.RemoteAddr
	ips := strings.Split(ip, ":")
	username := ips[0]

	member := logic.GetMember(username)
	if member != nil {
		c.Set("name", member.GetName())
		c.Set("username", username)
		return
	}

	// register
	member = &models.Member{
		Name:     username,
		Username: username,
	}
	err := database.Db.AddMembers(*member)
	if err != nil {
		return
	}

	c.Set("name", member.GetName())
	c.Set("username", member.GetUsername())

	// name := ""
	// switch ip := c.Request.RemoteAddr[:13]; ip {
	// case "192.168.0.106":
	// 	name = "奇異"
	// case "192.168.0.121":
	// 	name = "QQ詩"
	// case "192.168.0.135":
	// 	name = "廷"
	// case "192.168.0.108":
	// 	name = "叡"
	// case "192.168.0.137":
	// 	name = "雞排"
	// case "192.168.0.115":
	// 	name = "傑哥"
	// case "192.168.0.144":
	// 	name = "哥"
	// case "192.168.0.122":
	// 	name = "宏哥"
	// case "192.168.0.104":
	// 	name = "小麥"
	// case "192.168.0.123":
	// 	name = "hank"
	// case "192.168.0.136":
	// 	name = "Indy"
	// case "192.168.0.128":
	// 	name = "尾"
	// default:
	// 	name = ip
	// }

}
