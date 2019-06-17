package menu

import (
	"net/http"
	"orderfood/src/handler/models/resp"
	"orderfood/src/logic"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddShop 新增商店
// @Tags menu
// @Summary 新增商店
// @Description 新增商店
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param name formData string true "商名"
// @Success 200 {object} resp.Shop "菜單"
// @Failure 500 {string} string "内部错误"
// @Router /menu/shop [post]
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

// GetShop 取得商店
// @Tags menu
// @Summary 取得商店
// @Description 取得商店
// @Produce  json
// @Param id query string false "編號"
// @Param name query string false "商名"
// @Success 200 {array} resp.Shop "菜單"
// @Failure 500 {string} string "内部错误"
// @Router /menu/shop [get]
func GetShop(c *gin.Context) {
	shopIDStr := c.Query("id")
	shopID, err := strconv.Atoi(shopIDStr)
	if err != nil {
		shopID = 0
	}
	shopName := c.Query("name")

	data, err := logic.GetShop(shopID, shopName)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result := make([]resp.Shop, 0)
	for _, v := range data {
		result = append(result, resp.Shop{
			ID:   int32(v.GetID()),
			Name: v.GetName(),
		})
	}
	c.JSON(http.StatusOK, result)
}

// UpdateShop 修改商店
// @Tags menu
// @Summary 修改商店
// @Description 修改商店
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param id path int true "編號"
// @Param name formData string true "店名"
// @Success 200 {string} string "結果"
// @Failure 500 {string} string "内部错误"
// @Router /menu/shop/{id} [put]
func UpdateShop(c *gin.Context) {
	shopIDStr := c.Param("id")
	shopID, err := strconv.Atoi(shopIDStr)
	if err != nil {
		shopID = 0
	}
	shopName := c.PostForm("name")

	success, err := logic.UpdateShop(shopID, shopName)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, success)
}

// DeleteShop 刪除商店
// @Tags menu
// @Summary 刪除商店
// @Description 刪除商店
// @Produce  json
// @Param id path int true "編號"
// @Success 200 {string} result "成功"
// @Failure 500 {string} string "内部错误"
// @Router /menu/shop/{id} [delete]
func DeleteShop(c *gin.Context) {
	shopIDStr := c.Param("id")
	shopID, err := strconv.Atoi(shopIDStr)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	success, err := logic.DeleteShop(shopID)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, success)
}
