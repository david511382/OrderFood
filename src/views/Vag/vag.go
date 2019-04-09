package vag

import (
	"net/http"
	views "orderfood/src/templates"

	"github.com/gin-gonic/gin"
)

var MenuData = []views.MenuKind{
	views.MenuKind{
		Items: []views.MenuItem{
			views.MenuItem{
				Name:   "炒麵",
				Prices: []int{15, 0},
			},
			views.MenuItem{
				Name:   "烏龍麵",
				Prices: []int{15, 0},
			},
			views.MenuItem{
				Name:   "米粉",
				Prices: []int{20, 0},
			},
		},
		Size:              []string{"小", "大"},
		RequiredSelection: nil,
		CheckOption: []views.KindOption{
			views.KindOption{
				Name:  "加甜辣醬",
				Price: 0},
			views.KindOption{
				Name:  "加辣油",
				Price: 0},
		},
	},
	views.MenuKind{
		Items: []views.MenuItem{
			views.MenuItem{
				Name:   "南投意麵",
				Prices: []int{30},
			},
			views.MenuItem{
				Name:   "御膳湯麵",
				Prices: []int{30},
			},
			views.MenuItem{
				Name:   "當歸麵線",
				Prices: []int{40},
			},
			views.MenuItem{
				Name:   "切仔麵",
				Prices: []int{30},
			},
			views.MenuItem{
				Name:   "麻醬麵",
				Prices: []int{30},
			},
			views.MenuItem{
				Name:   "什錦麵",
				Prices: []int{40},
			},
			views.MenuItem{
				Name:   "米粉湯",
				Prices: []int{30},
			},
			views.MenuItem{
				Name:   "冬粉湯",
				Prices: []int{30},
			},
			views.MenuItem{
				Name:   "香菇羹飯",
				Prices: []int{40},
			},
			views.MenuItem{
				Name:   "香菇羹麵",
				Prices: []int{40},
			},
			views.MenuItem{
				Name:   "香菇羹米粉",
				Prices: []int{40},
			},
			views.MenuItem{
				Name:   "香樁炒飯",
				Prices: []int{20},
			},
			views.MenuItem{
				Name:   "素香飯",
				Prices: []int{20},
			},
			views.MenuItem{
				Name:   "苦瓜飯",
				Prices: []int{20},
			},

			views.MenuItem{
				Name:   "大豆干",
				Prices: []int{15},
			},
			views.MenuItem{
				Name:   "豆包",
				Prices: []int{10},
			},
			views.MenuItem{
				Name:   "包心丸(2個)",
				Prices: []int{15},
			},
			views.MenuItem{
				Name:   "小豆干",
				Prices: []int{10},
			},
			views.MenuItem{
				Name:   "海帶",
				Prices: []int{10},
			},
			views.MenuItem{
				Name:   "素肚",
				Prices: []int{10},
			},
			views.MenuItem{
				Name:   "素雞",
				Prices: []int{10},
			},
			views.MenuItem{
				Name:   "菜捲",
				Prices: []int{15},
			},
			views.MenuItem{
				Name:   "蘭花干",
				Prices: []int{15},
			},
			views.MenuItem{
				Name:   "芋頭簽",
				Prices: []int{10},
			},
			views.MenuItem{
				Name:   "火腿",
				Prices: []int{10},
			},
			views.MenuItem{
				Name:   "百頁豆腐",
				Prices: []int{15},
			},
			views.MenuItem{
				Name:   "米血",
				Prices: []int{10},
			},
			views.MenuItem{
				Name:   "苦瓜",
				Prices: []int{10},
			},
			views.MenuItem{
				Name:   "蘿蔔",
				Prices: []int{15},
			},
			views.MenuItem{
				Name:   "蒟蒻",
				Prices: []int{5},
			},
			views.MenuItem{
				Name:   "燙青菜",
				Prices: []int{20},
			},
		},
		Size:              make([]string, 0),
		RequiredSelection: make([]map[string]int, 0),
		CheckOption: []views.KindOption{
			views.KindOption{
				Name:  "加甜辣醬",
				Price: 0},
			views.KindOption{
				Name:  "加辣油",
				Price: 0},
		},
	},
	views.MenuKind{
		Items: []views.MenuItem{
			views.MenuItem{
				Name:   "四物湯",
				Prices: []int{20},
			},
			views.MenuItem{
				Name:   "香菇羹湯",
				Prices: []int{20},
			},
			views.MenuItem{
				Name:   "紫菜湯",
				Prices: []int{15},
			},
			views.MenuItem{
				Name:   "豆腐湯",
				Prices: []int{15},
			},
			views.MenuItem{
				Name:   "豆皮湯",
				Prices: []int{15},
			},
			views.MenuItem{
				Name:   "油豆腐湯",
				Prices: []int{15},
			},
			views.MenuItem{
				Name:   "綜合湯",
				Prices: []int{15},
			},
			views.MenuItem{
				Name:   "菜捲湯",
				Prices: []int{20},
			},

			views.MenuItem{
				Name:   "仙草蜜",
				Prices: []int{10},
			},
			views.MenuItem{
				Name:   "豆漿",
				Prices: []int{10},
			},
			views.MenuItem{
				Name:   "紅茶",
				Prices: []int{10},
			},
			views.MenuItem{
				Name:   "洛神花茶",
				Prices: []int{10},
			},
		},
		Size:              make([]string, 0),
		RequiredSelection: make([]map[string]int, 0),
		CheckOption:       make([]views.KindOption, 0),
	},
}

const plus = 5

func init() {
	arr := MenuData[0].Items
	for i := 0; i < len(arr); i++ {
		arr[i].Prices[1] = arr[i].Prices[0] + plus
	}
}

func View(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)

	views.Order(
		c,
		imgHTML,
		views.NameHTML,
		views.FormHeadHTML,
		views.MenuHTML,
		views.ShopcartHTML,
		views.FormEndHTML,
		views.ResultHTML,
	)
}

func imgHTML() string {
	return `<img src="/src/img/vag.jpg" alt="Smiley face">`
}
