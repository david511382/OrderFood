package menu

import (
	"net/http"
	"orderfood/src/handler/models/resp"
	"orderfood/src/logic"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddItem 新增商品
// @Tags menu
// @Summary 新增商品
// @Description 新增商品
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param shopID formData int true "商店"
// @Param name formData string true "商名"
// @Param price formData int false "價格"
// @Success 200 {object} resp.Item "菜單"
// @Failure 500 {string} string "内部错误"
// @Router /menu/item [post]
func AddItem(c *gin.Context) {
	shopIDStr := c.PostForm("shopID")
	shopID, err := strconv.Atoi(shopIDStr)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, nil)
		return
	}
	itemName := c.PostForm("name")
	if itemName == "" {
		c.AbortWithError(http.StatusBadRequest, nil)
		return
	}
	priceStr := c.PostForm("price")
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		price = 0
	}

	data, err := logic.AddItem(shopID, itemName, price)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result := &resp.Item{
		ID:      int32(data.GetID()),
		Name:    data.GetName(),
		Price:   int32(data.GetPrice()),
		Options: "",
	}

	c.JSON(http.StatusOK, result)
}

// GetItem 取得商品
// @Tags menu
// @Summary 取得商品
// @Description 取得商品
// @Produce  json
// @Param shopID path int true "商店編號"
// @Param optionID query int false "選單編號"
// @Success 200 {array} resp.Item "菜單"
// @Failure 500 {string} string "内部错误"
// @Router /menu/item/{shopID} [get]
func GetItem(c *gin.Context) {
	shopIDStr := c.Param("shopID")
	shopID, err := strconv.Atoi(shopIDStr)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	optionIDStr := c.Query("optionID")
	optionID, err := strconv.Atoi(optionIDStr)
	if err != nil {
		optionID = 0
	}

	data, err := logic.GetItem(shopID, optionID)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result := make([]resp.Item, 0)
	for _, v := range data {
		result = append(result, resp.Item{
			ID:      int32(v.GetID()),
			Name:    v.GetName(),
			Price:   int32(v.GetPrice()),
			Options: "",
		})
	}
	c.JSON(http.StatusOK, result)
}

// UpdateItem 修改商品
// @Tags menu
// @Summary 修改商品
// @Description 修改商品
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param id formData int true "編號"
// @Param name formData string false "品名"
// @Param price formData int false "價格"
// @Success 200 {string} string "結果"
// @Failure 500 {string} string "内部错误"
// @Router /menu/item [put]
func UpdateItem(c *gin.Context) {
	itemIDStr := c.PostForm("id")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	itemName := c.PostForm("name")
	priceStr := c.PostForm("price")
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		price = 0
		if itemName == "" {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	success, err := logic.UpdateItem(itemID, itemName, price)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, success)
}

// DeleteItem 刪除商品
// @Tags menu
// @Summary 刪除商品
// @Description 刪除商品
// @Produce  json
// @Param id path int true "編號"
// @Success 200 {string} result "成功"
// @Failure 500 {string} string "内部错误"
// @Router /menu/item/{id} [delete]
func DeleteItem(c *gin.Context) {
	itemIDStr := c.Param("id")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	success, err := logic.DeleteItem(itemID)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, success)
}
