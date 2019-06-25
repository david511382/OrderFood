package menu

import (
	"net/http"
	"orderfood/src/handler/models/resp"
	managerLgc "orderfood/src/logic/manager"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddOption 新增商品選單
// @Tags menu
// @Summary 新增商品選單
// @Description 新增商品選單
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param selectionName formData string true "選單選項"
// @Param selectNum formData int false "必選數"
// @Success 200 {object} resp.Option "商品選單"
// @Failure 500 {string} string "内部错误"
// @Router /menu/option [post]
func AddOption(c *gin.Context) {
	selectionName := c.PostForm("selectionName")
	if selectionName == "" {
		c.AbortWithError(http.StatusBadRequest, nil)
		return
	}
	selectNumStr := c.PostForm("selectNum")
	selectNum, err := strconv.Atoi(selectNumStr)
	if err != nil {
		selectNum = 0
	}

	data, err := managerLgc.AddOption(selectNum,selectionName)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result := &resp.Option{
		ID:        int32(data.GetID()),
		SelectNum: int32(data.GetSelect_Num()),
	}
	c.JSON(http.StatusOK, result)
}

// UpdateOption 修改商品選單
// @Tags menu
// @Summary 修改商品選單
// @Description 修改商品選單
// @Produce  json
// @Param selectNum formData int true "商品選單"
// @Success 200 {string} string "成功"
// @Failure 500 {string} string "内部错误"
// @Router /menu/option/{id} [put]
func UpdateOption(c *gin.Context) {
	selectNumStr := c.Param("id")
	selectNum, err := strconv.Atoi(selectNumStr)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	} else if selectNum < 0 {
		c.AbortWithError(http.StatusBadRequest, nil)
		return
	}
	idStr := c.PostForm("selectNum")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	} else if id <= 0 {
		c.AbortWithError(http.StatusBadRequest, nil)
		return
	}

	success, err := managerLgc.UpdateOption(id, selectNum)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, success)
}

// DeleteOption 刪除商品選單
// @Tags menu
// @Summary 刪除商品選單
// @Description 刪除商品選單
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {string} result "成功"
// @Failure 500 {string} string "内部错误"
// @Router /menu/option/{id} [delete]
func DeleteOption(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	} else if id <= 0 {
		c.AbortWithError(http.StatusBadRequest, nil)
		return
	}

	success, err := managerLgc.DeleteOption(id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, success)
}
