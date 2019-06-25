package manager

import (
	"github.com/gin-gonic/gin"
)

// AddOption 建立新選單
// @Tags manager
// @Summary 建立新選單
// @Description 建立新選單
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param data formData MenuOption true "新選單"
// @Success 200 {string} string "商店"
// @Failure 500 {string} string "内部错误"
// @Security ApiKeyAuth
// @Router /manager/newoption [post]
func AddOption(c *gin.Context) {
	// view := c.PostForm("view")

	// logic.SetView(view)

	// c.String(http.StatusOK, view)
}
