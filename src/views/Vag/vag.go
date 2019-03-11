package vag

import (
	"net/http"
	views "orderfood/src/templates"

	"github.com/gin-gonic/gin"
)

type vagItem struct {
	Name  string
	Price int
}

var MenuData = [...]vagItem{
	vagItem{
		Name:  "炒麵",
		Price: 35,
	},
	vagItem{
		Name:  "烏龍麵",
		Price: 35,
	},
}

func View(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)

	views.Order(
		c,
		imgHTML,
		views.NameHTML,
		views.FormHeadHTML,
		views.MenuHTML,
		shopcartHTML,
		views.FormEndHTML,
		views.ResultHTML,
	)
}

func imgHTML() string {
	return `<img src="/src/img/vag.jpg" alt="Smiley face" height="400" width="700">`
}

func shopcartHTML() string {
	return `
		</br>
		購物清單:
		<div id="shopcart"></div>
		</br>
		總計:<font id="total"></font> 元
		</br>
		<script src="/src/js/vag.js"></script>`
}
