package menu

import (
	"net/http"
	"orderfood/src/handler/models/resp"
	managerLgc "orderfood/src/logic/manager"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddSelection 新增選單選項
// @Tags menu
// @Summary 新增選單選項
// @Description 新增選單選項
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param optionID formData int true "選單編號"
// @Param name formData string true "名稱"
// @Param price formData int false "價格"
// @Success 200 {object} resp.MenuSelection "菜單"
// @Failure 500 {string} string "内部错误"
// @Router /menu/selection [post]
func AddSelection(c *gin.Context) {
	optionIDStr := c.PostForm("optionID")
	optionID, err := strconv.Atoi(optionIDStr)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	} else if optionID <= 0 {
		c.AbortWithError(http.StatusBadRequest, nil)
		return
	}
	name := c.PostForm("name")
	if name == "" {
		c.AbortWithError(http.StatusBadRequest, nil)
		return
	}
	priceStr := c.PostForm("price")
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		price = 0
	}

	data, err := managerLgc.AddSelection(optionID, price, name)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result := &resp.MenuSelection{
		ID:    int32(data.GetID()),
		Name:  data.GetName(),
		Price: int32(data.GetPrice()),
	}
	c.JSON(http.StatusOK, result)
}

// UpdateSelection 修改選單選項
// @Tags menu
// @Summary 修改選單選項
// @Description 修改選單選項
// @Produce  json
// @Param name formData string false "名稱"
// @Param price formData int false "價格"
// @Success 200 {string} string "結果"
// @Failure 500 {string} string "内部错误"
// @Router /menu/selection/{id} [put]
func UpdateSelection(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	} else if id <= 0 {
		c.AbortWithError(http.StatusBadRequest, nil)
		return
	}
	name := c.PostForm("name")
	priceStr := c.PostForm("price")
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		if name == "" {
			c.AbortWithError(http.StatusBadRequest, nil)
			return
		}
		price = -1
	}

	success, err := managerLgc.UpdateSelection(id, price, name)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, success)
}

// DeleteSelection 刪除選單選項
// @Tags menu
// @Summary 刪除選單選項
// @Description 刪除選單選項
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {string} result "成功"
// @Failure 500 {string} string "内部错误"
// @Router /menu/selection/{id} [get]
func DeleteSelection(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	} else if id <= 0 {
		c.AbortWithError(http.StatusBadRequest, nil)
		return
	}

	success, err := managerLgc.DeleteSelection(id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, success)
}
