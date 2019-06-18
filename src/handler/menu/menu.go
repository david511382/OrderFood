package menu

import (
	"net/http"
	"orderfood/src/logic"

	"github.com/gin-gonic/gin"
)

// GetMenu 取得菜單
// @Tags menu
// @Summary 取得菜單
// @Description 取得菜單
// @Produce  json
// @Param shop path string true "商店"
// @Success 200 {array} resp.ShopMenu "菜單"
// @Failure 500 {string} string "内部错误"
// @Router /menu/menu/{shop} [get]
func GetMenu(c *gin.Context) {
	shop := c.Param("shop")

	data, err := logic.GetMenu(shop)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, data)
}

// GetShopMenu 取得菜單
// @Tags menu
// @Summary 取得菜單
// @Description 取得菜單
// @Produce  json
// @Success 200 {array} resp.ShopMenu "菜單"
// @Failure 500 {string} string "内部错误"
// @Router /menu/shopmenu [get]
func GetShopMenu(c *gin.Context) {
	data, err := logic.GetShopMenu()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, data)
}
