package auth

import (
	"net/http"
	"orderfood/src/database/models"
	"orderfood/src/logic"

	"github.com/gin-gonic/gin"
)

// Register 註冊
// @Tags auth
// @Summary 註冊
// @Description 註冊
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param name formData string true "稱號"
// @Param username formData string true "帳號"
// @Param password formData string true "密碼"
// @Success 200 {string} string ""
// @Failure 500 {string} string "内部错误"
// @Router /auth/register [post]
func Register(c *gin.Context) {
	name := c.PostForm("name")
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" {
		c.AbortWithError(http.StatusBadRequest, nil)
		return
	}

	member := &models.Member{
		Name:     name,
		Username: username,
		Password: password,
	}
	err := logic.Register(member)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
}
