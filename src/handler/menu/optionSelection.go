package menu

import (
	"github.com/gin-gonic/gin"
)

// AddOptionSelection 新增商店
// @Tags menu
// @Summary 新增商店
// @Description 新增商店
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param name formData string true "商店"
// @Success 200 {object} resp.MenuOption "菜單"
// @Failure 500 {string} string "内部错误"
// @Router /menu/optionSelection [post]
func AddOptionSelection(c *gin.Context) {
	// shopName := c.PostForm("name")
	// if shopName == "" {
	// 	c.AbortWithError(http.StatusBadRequest, nil)
	// 	return
	// }

	// data, err := logic.AddOptionSelection(shopName)
	// if err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	// c.JSON(http.StatusOK, data)
}

// GetOptionSelection 取得商店
// @Tags menu
// @Summary 取得商店
// @Description 取得商店
// @Produce  json
// @Success 200 {array} resp.MenuOption "菜單"
// @Failure 500 {string} string "内部错误"
// @Router /menu/optionSelection [get]
func GetOptionSelection(c *gin.Context) {
	// data, err := logic.GetOptionSelection()
	// if err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	// response := make([]resp.MenuOption, 0)
	// for _, v := range data {
	// 	response = append(response, resp.MenuOption{
	// 		ID:   v.GetID(),
	// 		Name: v.GetName(),
	// 	})
	// }
	// c.JSON(http.StatusOK, response)
}

// UpdateOptionSelection 修改商店
// @Tags menu
// @Summary 修改商店
// @Description 修改商店
// @Produce  json
// @Param name formData string true "商店"
// @Success 200 {array} resp.MenuOption "菜單"
// @Failure 500 {string} string "内部错误"
// @Router /menu/optionSelection [put]
func UpdateOptionSelection(c *gin.Context) {
	// data, err := logic.GetOptionSelection()
	// if err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	// response := make([]resp.MenuOption, 0)
	// for _, v := range data {
	// 	response = append(response, resp.MenuOption{
	// 		ID:   v.GetID(),
	// 		Name: v.GetName(),
	// 	})
	// }
	// c.JSON(http.StatusOK, response)
}

// DeleteOptionSelection 刪除商店
// @Tags menu
// @Summary 刪除商店
// @Description 刪除商店
// @Produce  json
// @Param id formData int true "ID"
// @Success 200 {string} result "成功"
// @Failure 500 {string} string "内部错误"
// @Router /menu/optionSelection [get]
func DeleteOptionSelection(c *gin.Context) {
	// data, err := logic.GetOptionSelection()
	// if err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	// response := make([]resp.MenuOption, 0)
	// for _, v := range data {
	// 	response = append(response, resp.MenuOption{
	// 		ID:   v.GetID(),
	// 		Name: v.GetName(),
	// 	})
	// }
	// c.JSON(http.StatusOK, response)
}