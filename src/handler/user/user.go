package user

import (
	"fmt"
	"net/http"
	"orderfood/src/handler/ws"
	"orderfood/src/logic"
	rice "orderfood/src/views/Rice"
	vag "orderfood/src/views/Vag"
	"strings"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	switch logic.GetView() {
	case logic.R:
		rice.View(c)
	case logic.V:
		vag.View(c)
	}
}

func GetMenu(c *gin.Context) {
	switch logic.GetView() {
	case logic.R:
		c.JSON(http.StatusOK, rice.MenuData)
	case logic.V:
		c.JSON(http.StatusOK, vag.MenuData)
	}
}

func Order(c *gin.Context) {
	orderStr := c.PostForm("orders")

	v, ok := c.Get("username")
	if !ok{

	}
	name := v.(string)

	orders := strings.Split(orderStr, "\n")
	orders = orders[:len(orders)-1]
	logic.UserOrders[name] = orders

	log := fmt.Sprintf(
		"%s\n%s\n",
		name, orderStr,
	)
	fmt.Print(log)

	totalList := logic.IntegrationOrders()
	c.String(http.StatusOK, totalList)

	ws.Notify(totalList)
}

func UserOrder(c *gin.Context) {
	result := ""
	for name, orders := range logic.UserOrders {
		result += name + "\n"
		for _, order := range orders {
			result += order + "\n"
		}
		result += "\n"
	}

	c.String(http.StatusOK, result)
}
