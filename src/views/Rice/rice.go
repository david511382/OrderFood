package rice

import (
	"net/http"
	views "orderfood/src/templates"

	"github.com/gin-gonic/gin"
)

type menuItem struct {
	Name   string
	Half   int
	Normal int
}

var MenuData = [...]menuItem{
	menuItem{
		Name:   "蔬果沙拉",
		Half:   35,
		Normal: 65,
	},
	menuItem{
		Name:   "蔬果雞腿沙拉",
		Half:   35,
		Normal: 65,
	},
	menuItem{
		Name:   "洋蔥牛肉",
		Half:   40,
		Normal: 70,
	},
	menuItem{
		Name:   "鳳梨牛肉",
		Half:   45,
		Normal: 75,
	},
	menuItem{
		Name:   "黑胡椒燒肉",
		Half:   35,
		Normal: 55,
	},
	menuItem{
		Name:   "洋蔥燒肉",
		Half:   35,
		Normal: 55,
	},
	menuItem{
		Name:   "香蒜培根",
		Half:   30,
		Normal: 50,
	},
	menuItem{
		Name:   "辣子雞丁",
		Half:   30,
		Normal: 50,
	},
	menuItem{
		Name:   "鳳梨雞丁",
		Half:   35,
		Normal: 55,
	},
	menuItem{
		Name:   "蘋果雞丁",
		Half:   35,
		Normal: 55,
	},
	menuItem{
		Name:   "蘋果雞腿",
		Half:   35,
		Normal: 55,
	},
	menuItem{
		Name:   "蘋果招牌",
		Half:   35,
		Normal: 60,
	},
	menuItem{
		Name:   "肉鬆沙拉",
		Half:   30,
		Normal: 50,
	},
	menuItem{
		Name:   "鮪魚沙拉",
		Half:   30,
		Normal: 50,
	},
	menuItem{
		Name:   "鮪魚肉鬆",
		Half:   35,
		Normal: 55,
	},
	menuItem{
		Name:   "糖醋雞腿",
		Half:   35,
		Normal: 55,
	},
	menuItem{
		Name:   "凱薩雞腿沙拉",
		Half:   30,
		Normal: 50,
	},
	menuItem{
		Name:   "凱薩雞腿肉鬆",
		Half:   35,
		Normal: 55,
	},
	menuItem{
		Name:   "招牌綜合",
		Half:   35,
		Normal: 55,
	},
	menuItem{
		Name:   "泡菜招牌",
		Half:   35,
		Normal: 60,
	},
	menuItem{
		Name:   "泡菜雞丁",
		Half:   35,
		Normal: 55,
	},
	menuItem{
		Name:   "泡菜培根",
		Half:   35,
		Normal: 55,
	},
	menuItem{
		Name:   "泡菜燒肉",
		Half:   35,
		Normal: 55,
	},
	menuItem{
		Name:   "泡菜牛肉",
		Half:   45,
		Normal: 75,
	},
	menuItem{
		Name:   "泡菜雞腿",
		Half:   35,
		Normal: 55,
	},
	menuItem{
		Name:   "香菇素鬆",
		Half:   30,
		Normal: 50,
	},
	menuItem{
		Name:   "海苔香酥",
		Half:   30,
		Normal: 50,
	},
	menuItem{
		Name:   "素蘋果沙拉",
		Half:   35,
		Normal: 65,
	},
	menuItem{
		Name:   "素蘋果招牌",
		Half:   35,
		Normal: 60,
	},
	menuItem{
		Name:   "招牌素總匯",
		Half:   35,
		Normal: 55,
	},
	menuItem{
		Name:   "牛蒡豆絲火腿",
		Half:   30,
		Normal: 50,
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
	return `<img src="/src/img/rice.jpg" alt="Smiley face" height="400" width="700">`
}

func shopcartHTML() string {
	return `
		</br>
		購物清單:
		<div id="shopcart"></div>
		</br>
		總計:<font id="total"></font> 元
		</br>
		<script src="/src/js/rice.js"></script>`
}
