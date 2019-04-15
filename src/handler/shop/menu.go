package shop

import (
	"net/http"
	"orderfood/src/logic"

	"github.com/gin-gonic/gin"
)

// GetMenu 取得菜單
// @Tags shop
// @Summary 取得菜單
// @Description 取得菜單
// @Produce  json
// @Success 200 {array} resp.MenuKind "菜單"
// @Failure 500 {string} string "内部错误"
// @Router /shop/menu [get]
func GetMenu(c *gin.Context) {
	data, err := logic.GetMenu()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, data)
}
