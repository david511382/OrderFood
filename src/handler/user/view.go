package user

import (
	"net/http"
	"orderfood/src/logic"

	rice "orderfood/src/views/Rice"
	vag "orderfood/src/views/Vag"

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

// GetMenu 取得菜單
// @Tags shop
// @Summary 取得菜單
// @Description 取得菜單
// @Produce  json
// @Success 200 {string} string "菜單"
// @Failure 500 {string} string "内部错误"
// @Router /menu [get]
func GetMenu(c *gin.Context) {
	switch logic.GetView() {
	case logic.R:
		c.JSON(http.StatusOK, rice.MenuData)
	case logic.V:
		c.JSON(http.StatusOK, vag.MenuData)
	}
}
