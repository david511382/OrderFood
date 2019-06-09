package menu

 import (
// 	"net/http"
// 	"orderfood/src/logic"
// 	"strconv"

// 	"orderfood/src/handler/models/resp"

	"github.com/gin-gonic/gin"
)

// AddShop 新增商店
// @Tags menu
// @Summary 新增商店
// @Description 新增商店
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param name formData string true "商店"
// @Success 200 {object} resp.Shop "菜單"
// @Failure 500 {string} string "内部错误"
// @Router /menu/shop [post]
func AddShop(c *gin.Context) {
	// shopName := c.PostForm("name")
	// if shopName == "" {
	// 	c.AbortWithError(http.StatusBadRequest, nil)
	// 	return
	// }

	// data, err := logic.AddShop(shopName)
	// if err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	// c.JSON(http.StatusOK, data)
}

// GetShop 取得商店
// @Tags menu
// @Summary 取得商店
// @Description 取得商店
// @Produce  json
// @Success 200 {array} resp.Shop "菜單"
// @Failure 500 {string} string "内部错误"
// @Router /menu/shop [get]
func GetShop(c *gin.Context) {
	// data, err := logic.GetShop()
	// if err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	// response := make([]resp.Shop, 0)
	// for _, v := range data {
	// 	response = append(response, resp.Shop{
	// 		ID:   v.GetID(),
	// 		Name: v.GetName(),
	// 	})
	// }
	// c.JSON(http.StatusOK, response)
}

// UpdateShop 修改商店
// @Tags menu
// @Summary 修改商店
// @Description 修改商店
// @Produce  json
// @Param name formData string true "商店"
// @Success 200 {array} resp.Shop "菜單"
// @Failure 500 {string} string "内部错误"
// @Router /menu/shop [put]
func UpdateShop(c *gin.Context) {
	// data, err := logic.GetShop()
	// if err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	// response := make([]resp.Shop, 0)
	// for _, v := range data {
	// 	response = append(response, resp.Shop{
	// 		ID:   v.GetID(),
	// 		Name: v.GetName(),
	// 	})
	// }
	// c.JSON(http.StatusOK, response)
}

// DeleteShop 刪除商店
// @Tags menu
// @Summary 刪除商店
// @Description 刪除商店
// @Produce  json
// @Param id formData int true "ID"
// @Success 200 {string} result "成功"
// @Failure 500 {string} string "内部错误"
// @Router /menu/shop [get]
func DeleteShop(c *gin.Context) {
	// data, err := logic.GetShop()
	// if err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	// response := make([]resp.Shop, 0)
	// for _, v := range data {
	// 	response = append(response, resp.Shop{
	// 		ID:   v.GetID(),
	// 		Name: v.GetName(),
	// 	})
	// }
	// c.JSON(http.StatusOK, response)
}