package manager

import (
	"net/http"
	"orderfood/src/logic"

	"github.com/gin-gonic/gin"
)

// ChangeView 更改商店
// @Tags shop
// @Summary 更改商店
// @Description 更改商店
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param view formData string true "商店"
// @Success 200 {string} string "商店"
// @Failure 500 {string} string "内部错误"
// @Security ApiKeyAuth
// @Router /shop [put]
func ChangeView(c *gin.Context) {
	view := c.PostForm("view")

	logic.SetView(view)

	c.String(http.StatusOK, view)
}
