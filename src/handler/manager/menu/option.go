package menu

import (
	"net/http"
	managerLgc "orderfood/src/logic/manager"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UpdateOption 修改商品選單
// @Tags menu
// @Summary 修改商品選單
// @Description 修改商品選單
// @Produce  json
// @Param selectNum formData int true "商品選單"
// @Success 200 {string} string "成功"
// @Failure 500 {string} string "内部错误"
// @Router /manager/menu/option/{id} [put]
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
// @Router /manager/menu/option/{id} [delete]
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
