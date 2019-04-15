package shop

import (
	"net/http"
	"orderfood/src/logic"

	"github.com/gin-gonic/gin"
)

// AddShop 新增商店
// @Tags shop
// @Summary 新增商店
// @Description 新增商店
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param name formData string true "商店"
// @Success 200 {array} resp.Shop "菜單"
// @Failure 500 {string} string "内部错误"
// @Router /shop [post]
func AddShop(c *gin.Context) {
	shopName := c.PostForm("name")
	if shopName == "" {
		c.AbortWithError(http.StatusBadRequest, nil)
		return
	}

	data, err := logic.AddShop(shopName)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, data)
}
