package menu

import (
	// "net/http"
	// "orderfood/src/handler/models/resp"
	// "orderfood/src/logic"
	// "strconv"

	"github.com/gin-gonic/gin"
)

// AddOption 新增商品選單
// @Tags menu
// @Summary 新增商品選單
// @Description 新增商品選單
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param shopID formData int true "商品選單"
// @Param leastSelectNum formData leastSelectNum false "商品選單"
// @Success 200 {object} resp.MenuOption "菜單"
// @Failure 500 {string} string "内部错误"
// @Router /menu/group [post]
func AddOption(c *gin.Context) {
	// shopIDStr := c.PostForm("shopID")
	// shopID, err := strconv.Atoi(shopIDStr)
	// if err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }
	// leastSelectNumStr := c.PostForm("leastSelectNum")
	// leastSelectNum, err := strconv.Atoi(leastSelectNumStr)
	// if err != nil {
	// 	leastSelectNum = 0
	// }

	// data, err := logic.AddOption(int32(shopID), int32(leastSelectNum))
	// if err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	// c.JSON(http.StatusOK, data)
}

// GetOption 取得商品選單
// @Tags menu
// @Summary 取得商品選單
// @Description 取得商品選單
// @Produce  json
// @Success 200 {array} resp.MenuOption "菜單"
// @Failure 500 {string} string "内部错误"
// @Router /menu/group/{shopID} [get]
func GetOption(c *gin.Context) {
	// shopIDStr := c.Param("shopID")
	// shopID, err := strconv.Atoi(shopIDStr)
	// if err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	// data, err := logic.GetOption(int32(shopID))
	// if err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	// response := make([]resp.MenuOption, 0)
	// for _, v := range data {
	// 	response = append(response, resp.MenuOption{
	// 		ID:   v.GetID(),
	// 	})
	// }
	// c.JSON(http.StatusOK, response)
}

// UpdateOption 修改商品選單
// @Tags menu
// @Summary 修改商品選單
// @Description 修改商品選單
// @Produce  json
// @Param name formData string true "商品選單"
// @Success 200 {array} resp.MenuOption "菜單"
// @Failure 500 {string} string "内部错误"
// @Router /menu/group [put]
func UpdateOption(c *gin.Context) {
	// data, err := logic.GetOption()
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

// DeleteOption 刪除商品選單
// @Tags menu
// @Summary 刪除商品選單
// @Description 刪除商品選單
// @Produce  json
// @Param id formData int true "ID"
// @Success 200 {string} result "成功"
// @Failure 500 {string} string "内部错误"
// @Router /menu/group [get]
func DeleteOption(c *gin.Context) {
	// data, err := logic.GetOption()
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
