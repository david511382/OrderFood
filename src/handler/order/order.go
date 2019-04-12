package order

import (
	"fmt"
	"net/http"
	"orderfood/src/handler/ws"
	"orderfood/src/logic"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetTotalOrders 取得所有訂單
// @Tags order
// @Summary 取得所有訂單
// @Description 取得所有訂單
// @Produce  json
// @Success 200 {string} string "餐點"
// @Failure 500 {string} string "内部错误"
// @Router /get/order [post]
func GetTotalOrders(c *gin.Context) {
	totalList := logic.IntegrationOrders()

	c.String(http.StatusOK, totalList)
}

// Order 訂餐
// @Tags order
// @Summary 訂餐
// @Description 訂餐
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param orders formData string true "餐點"
// @Success 200 {string} string "餐點"
// @Failure 500 {string} string "内部错误"
// @Security ApiKeyAuth
// @Router /post/order [post]
func Order(c *gin.Context) {
	orderStr := c.PostForm("orders")

	v, ok := c.Get("username")
	if !ok {

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

// UserOrder 取得訂單
// @Tags order
// @Summary 取得訂單
// @Description 取得訂單
// @Produce  json
// @Success 200 {string} string "餐點"
// @Failure 500 {string} string "内部错误"
// @Router /get/user/orders [post]
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
