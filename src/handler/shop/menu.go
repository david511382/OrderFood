package shop

import (
	"net/http"
	"orderfood/src/logic"

	"orderfood/src/handler/models/resp"

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

// AddItem 新增商品
// @Tags shop
// @Summary 新增商品
// @Description 新增商品
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param name formData string true "商品"
// @Success 200 {object} resp.MenuItem "商品"
// @Failure 500 {string} string "内部错误"
// @Router /shop/item/ [post]
func AddItem(c *gin.Context) {
	itemName := c.PostForm("name")
	if itemName == "" {
		c.AbortWithError(http.StatusBadRequest, nil)
		return
	}

	data, err := logic.AddItem(itemName)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, data)
}

// GetItem 取得商品
// @Tags shop
// @Summary 取得商品
// @Description 取得商品
// @Produce  json
// @Success 200 {array} resp.MenuItem "商品"
// @Failure 500 {string} string "内部错误"
// @Router /shop/item/ [get]
func GetItem(c *gin.Context) {
	data, err := logic.GetItem()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response := make([]*resp.MenuItem, 0)
	for _, v := range data {
		response = append(response, &resp.MenuItem{
			Name: v.GetName(),
		})
	}
	c.JSON(http.StatusOK, response)
}
