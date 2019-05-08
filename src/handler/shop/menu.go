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
// @Router /shop/item [post]
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
// @Router /shop/item [get]
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

// AddSize 新增尺寸
// @Tags shop
// @Summary 新增尺寸
// @Description 新增尺寸
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param name formData string true "尺寸"
// @Success 200 {object} resp.Size "尺寸"
// @Failure 500 {string} string "内部错误"
// @Router /shop/size [post]
func AddSize(c *gin.Context) {
	sizeName := c.PostForm("name")
	if sizeName == "" {
		c.AbortWithError(http.StatusBadRequest, nil)
		return
	}

	data, err := logic.AddSize(sizeName)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response := &resp.Size{
		ID:   data.GetID(),
		Size: data.GetName(),
	}
	c.JSON(http.StatusOK, response)
}

// GetSize 取得尺寸
// @Tags shop
// @Summary 取得尺寸
// @Description 取得尺寸
// @Produce  json
// @Success 200 {array} resp.Size "尺寸"
// @Failure 500 {string} string "内部错误"
// @Router /shop/size [get]
func GetSize(c *gin.Context) {
	data, err := logic.GetSize()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response := make([]*resp.Size, 0)
	for _, v := range data {
		response = append(response, &resp.Size{
			ID:   v.GetID(),
			Size: v.GetName(),
		})
	}
	c.JSON(http.StatusOK, response)
}

// AddKind 新增種類
// @Tags shop
// @Summary 新增種類
// @Description 新增種類
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param name formData string true "種類"
// @Success 200 {object} resp.KindOption "種類"
// @Failure 500 {string} string "内部错误"
// @Router /shop/kind [post]
func AddKind(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		c.AbortWithError(http.StatusBadRequest, nil)
		return
	}

	data, err := logic.AddKind(name)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response := &resp.KindOption{
		ID:   data.GetID(),
		Name: data.GetName(),
	}
	c.JSON(http.StatusOK, response)
}

// GetKind 取得種類
// @Tags shop
// @Summary 取得種類
// @Description 取得種類
// @Produce  json
// @Success 200 {array} resp.KindOption "種類"
// @Failure 500 {string} string "内部错误"
// @Router /shop/kind [get]
func GetKind(c *gin.Context) {
	data, err := logic.GetKind()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response := make([]*resp.KindOption, 0)
	for _, v := range data {
		response = append(response, &resp.KindOption{
			ID:   v.GetID(),
			Name: v.GetName(),
		})
	}
	c.JSON(http.StatusOK, response)
}
