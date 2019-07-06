package manager

import (
	"net/http"
	"orderfood/src/handler/models/reqs"
	managerLgc "orderfood/src/logic/manager"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

// AddOption 建立新選單
// @Tags manager
// @Summary 建立新選單
// @Description 建立新選單
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param menuOptionJS formData string true "新選單 JSON" default({"ShopID": 14,"SelectNum": 0,"Items":[{"Name":"i1","Price": 2}],"Selections":[{"Name": "s1","Price": 2}]})
// @Success 200 {object} resp.OptionMenu  "選單"
// @Failure 500 {string} string "内部错误"
// @Security ApiKeyAuth
// @Router /manager/menu/option [post]
func AddOption(c *gin.Context) {
	dataJS := c.PostForm("menuOptionJS")
	data := &reqs.MenuOption{}
	err := json.Unmarshal([]byte(dataJS), data)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	result, err := managerLgc.CreateOption(data)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
