package user

import (
	"net/http"
	"orderfood/src/database"
	"orderfood/src/handler/user/reqs"
	"orderfood/src/logic"

	"github.com/gin-gonic/gin"
)

func GetUserName(c *gin.Context) {
	v, ok := c.Get("name")
	if !ok {

	}
	name := v.(string)

	c.String(http.StatusOK, name)
}

func ModifyUser(c *gin.Context) {
	v, ok := c.Get("username")
	if !ok {
		c.AbortWithError(http.StatusBadRequest, nil)
		return
	}
	username := v.(string)

	req := reqs.ModifyUser{}
	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	member := logic.GetMember(username)
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
