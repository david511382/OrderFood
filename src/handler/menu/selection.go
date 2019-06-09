package menu

import (

	"github.com/gin-gonic/gin"
)

// AddSelection 新增商店
// @Tags menu
// @Summary 新增商店
// @Description 新增商店
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param name formData string true "商店"
// @Success 200 {object} resp.MenuSelection "菜單"
// @Failure 500 {string} string "内部错误"
// @Router /menu/selection [post]
func AddSelection(c *gin.Context) {
	// shopName := c.PostForm("name")
	// if shopName == "" {
	// 	c.AbortWithError(http.StatusBadRequest, nil)
	// 	return
	// }

	// data, err := logic.AddSelection(shopName)
	// if err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	// c.JSON(http.StatusOK, data)
}

// GetSelection 取得商店
// @Tags menu
// @Summary 取得商店
// @Description 取得商店
// @Produce  json
// @Success 200 {array} resp.MenuSelection "菜單"
// @Failure 500 {string} string "内部错误"
// @Router /menu/selection [get]
func GetSelection(c *gin.Context) {
	// data, err := logic.GetSelection()
	// if err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	// response := make([]resp.MenuSelection, 0)
	// for _, v := range data {
	// 	response = append(response, resp.MenuSelection{
	// 		ID:   v.GetID(),
	// 		Name: v.GetName(),
	// 	})
	// }
	// c.JSON(http.StatusOK, response)
}

// UpdateSelection 修改商店
// @Tags menu
// @Summary 修改商店
// @Description 修改商店
// @Produce  json
// @Param name formData string true "商店"
// @Success 200 {array} resp.MenuSelection "菜單"
// @Failure 500 {string} string "内部错误"
// @Router /menu/selection [put]
func UpdateSelection(c *gin.Context) {
	// data, err := logic.GetSelection()
	// if err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	// response := make([]resp.MenuSelection, 0)
	// for _, v := range data {
	// 	response = append(response, resp.MenuSelection{
	// 		ID:   v.GetID(),
	// 		Name: v.GetName(),
	// 	})
	// }
	// c.JSON(http.StatusOK, response)
}

// DeleteSelection 刪除商店
// @Tags menu
// @Summary 刪除商店
// @Description 刪除商店
// @Produce  json
// @Param id formData int true "ID"
// @Success 200 {string} result "成功"
// @Failure 500 {string} string "内部错误"
// @Router /menu/selection [get]
func DeleteSelection(c *gin.Context) {
	// data, err := logic.GetSelection()
	// if err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	// response := make([]resp.MenuSelection, 0)
	// for _, v := range data {
	// 	response = append(response, resp.MenuSelection{
	// 		ID:   v.GetID(),
	// 		Name: v.GetName(),
	// 	})
	// }
	// c.JSON(http.StatusOK, response)
}