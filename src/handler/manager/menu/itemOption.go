package menu

import (
	"net/http"
	"orderfood/src/handler/models/resp"
	managerLgc "orderfood/src/logic/manager"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddItemOption 商品加入選單
// @Tags menu
// @Summary 商品加入選單
// @Description 商品加入選單
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param itemID formData int true "商品編號"
// @Param optionID formData int true "選單編號"
// @Success 200 {object} resp.ItemOption "結果"
// @Failure 500 {string} string "内部错误"
// @Router /menu/itemoption [post]
func AddItemOption(c *gin.Context) {
	itemIDstr := c.PostForm("itemID")
	itemID, err := strconv.Atoi(itemIDstr)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	optionIDstr := c.PostForm("optionID")
	optionID, err := strconv.Atoi(optionIDstr)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data, err := managerLgc.AddItemOption(itemID, optionID)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result := &resp.ItemOption{
		ID:       int32(data.GetID()),
		ItemID:   int32(data.GetItem_ID()),
		OptionID: int32(data.GetOption_ID()),
	}

	c.JSON(http.StatusOK, result)
}

// DeleteItemOption 刪除選單的商品
// @Tags menu
// @Summary 刪除選單的商品
// @Description 刪除選單的商品
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {string} result "成功"
// @Failure 500 {string} string "内部错误"
// @Router /menu/itemoption/{id} [delete]
func DeleteItemOption(c *gin.Context) {
	itemOptionIDStr := c.Param("id")
	itemOptionID, err := strconv.Atoi(itemOptionIDStr)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	success, err := managerLgc.DeleteItemOption(itemOptionID)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, success)
}
