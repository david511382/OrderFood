package user

import (
	"net/http"
	"orderfood/src/database"
	"orderfood/src/handler/models/reqs"
	"orderfood/src/logic"

	"github.com/gin-gonic/gin"
)

// GetUserName 取得用戶名稱
// @Tags user
// @Summary 取得用戶名稱
// @Description 取得用戶名稱
// @Produce  json
// @Success 200 {string} string "用戶名稱"
// @Failure 500 {string} string "内部错误"
// @Security ApiKeyAuth
// @Router /user [get]
func GetUserName(c *gin.Context) {
	v, ok := c.Get("name")
	if !ok {
		c.AbortWithError(http.StatusBadRequest, nil)
		return
	}
	name := v.(string)

	c.String(http.StatusOK, name)
}

// ModifyUser 修改用戶資訊
// @Tags user
// @Summary 修改用戶資訊
// @Description 修改用戶名稱和密碼
// @Accept  json
// @Produce  json
// @Param json body reqs.ModifyUser true "修改用戶資訊"
// @Success 200 {object} models.Member "用戶資訊"
// @Failure 500 {string} string "内部错误"
// @Security ApiKeyAuth
// @Router /user [put]
func ModifyUser(c *gin.Context) {
	req := reqs.ModifyUser{}
	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	member := logic.GetMember(req.GetUsername())
	if member == nil {
		c.AbortWithError(http.StatusBadRequest, nil)
		return
	}
	if member.GetName() == req.GetName() && member.GetPassword() == req.GetPassword() {
		c.AbortWithError(http.StatusBadRequest, nil)
		return
	}

	member.Name = req.GetName()
	member.Password = req.GetPassword()

	err = database.Db.UpdateMembers(*member)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	logic.LoadMembers()

	c.JSON(http.StatusOK, member)
}
