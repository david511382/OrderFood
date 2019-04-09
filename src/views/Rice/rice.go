package rice

import (
	"net/http"
	views "orderfood/src/templates"

	"github.com/gin-gonic/gin"
)

var MenuData = []views.MenuKind{
	views.MenuKind{
		Items: []views.MenuItem{
			views.MenuItem{
				Name:   "蔬果沙拉",
				Prices: []int{35, 65, 0},
			},
			views.MenuItem{
				Name:   "蔬果雞腿沙拉",
				Prices: []int{35, 65, 75},
			},
			views.MenuItem{
				Name:   "洋蔥牛肉",
				Prices: []int{40, 70, 75},
			},

			views.MenuItem{
				Name:   "鳳梨牛肉",
				Prices: []int{45, 75, 75},
			},
			views.MenuItem{
				Name:   "黑胡椒燒肉",
				Prices: []int{35, 55, 75},
			},
			views.MenuItem{
				Name:   "洋蔥燒肉",
				Prices: []int{35, 55, 75},
			},
			views.MenuItem{
				Name:   "香蒜培根",
				Prices: []int{30, 50, 75},
			},
			views.MenuItem{
				Name:   "辣子雞丁",
				Prices: []int{30, 50, 75},
			},
			views.MenuItem{
				Name:   "鳳梨雞丁",
				Prices: []int{35, 55, 75},
			},
			views.MenuItem{
				Name:   "蘋果雞丁",
				Prices: []int{35, 55, 75},
			},
			views.MenuItem{
				Name:   "蘋果雞腿",
				Prices: []int{35, 55, 75},
			},
			views.MenuItem{
				Name:   "蘋果招牌",
				Prices: []int{35, 60, 75},
			},

			views.MenuItem{
				Name:   "肉鬆沙拉",
				Prices: []int{30, 50, 75},
			},
			views.MenuItem{
				Name:   "鮪魚沙拉",
				Prices: []int{30, 50, 75},
			},
			views.MenuItem{
				Name:   "鮪魚肉鬆",
				Prices: []int{35, 55, 75},
			},
			views.MenuItem{
				Name:   "糖醋雞腿",
				Prices: []int{35, 55, 75},
			},
			views.MenuItem{
				Name:   "凱薩雞腿沙拉",
				Prices: []int{30, 50, 75},
			},
			views.MenuItem{
				Name:   "凱薩雞腿肉鬆",
				Prices: []int{35, 55, 75},
			},
			views.MenuItem{
				Name:   "招牌綜合",
				Prices: []int{35, 55, 75},
			},
			views.MenuItem{
				Name:   "泡菜招牌",
				Prices: []int{35, 60, 75},
			},
			views.MenuItem{
				Name:   "泡菜雞丁",
				Prices: []int{35, 55, 75},
			},
			views.MenuItem{
				Name:   "泡菜培根",
				Prices: []int{35, 55, 75},
			},
			views.MenuItem{
				Name:   "泡菜燒肉",
				Prices: []int{35, 55, 75},
			},
			views.MenuItem{
				Name:   "泡菜牛肉",
				Prices: []int{45, 75, 75},
			},
			views.MenuItem{
				Name:   "泡菜雞腿",
				Prices: []int{35, 55, 75},
			},
			views.MenuItem{
				Name:   "香菇素鬆",
				Prices: []int{30, 50, 75},
			},
			views.MenuItem{
				Name:   "海苔香酥",
				Prices: []int{30, 50, 75},
			},
			views.MenuItem{
				Name:   "素蘋果沙拉",
				Prices: []int{35, 65, 75},
			},
			views.MenuItem{
				Name:   "素蘋果招牌",
				Prices: []int{35, 60, 75},
			},
			views.MenuItem{
				Name:   "招牌素總匯",
				Prices: []int{35, 55, 75},
			},
			views.MenuItem{
				Name:   "牛蒡豆絲火腿",
				Prices: []int{30, 50, 75},
			},
		},
		Size:              []string{"半捲", "正常", "加大"},
		RequiredSelection: nil,
		CheckOption: []views.KindOption{
			views.KindOption{
				Name:  "換拾穀米",
				Price: 5},
			views.KindOption{
				Name:  "加起司",
				Price: 10},
		},
	},
	views.MenuKind{
		Items: []views.MenuItem{
			views.MenuItem{
				Name:   "玉米濃湯",
				Prices: []int{30},
			},
			views.MenuItem{
				Name:   "昆布海苔湯",
				Prices: []int{30},
			},
			views.MenuItem{
				Name:   "特選紅茶",
				Prices: []int{25},
			},
			views.MenuItem{
				Name:   "鮮奶紅茶",
				Prices: []int{35},
			},
			views.MenuItem{
				Name:   "紫蘇梅汁",
				Prices: []int{30},
			},
			views.MenuItem{
				Name:   "冬瓜鮮奶茶",
				Prices: []int{35},
			},
			views.MenuItem{
				Name:   "古早味冬瓜茶",
				Prices: []int{25},
			},
			views.MenuItem{
				Name:   "冷泡烏龍綠茶",
				Prices: []int{25},
			},
		},
		Size:              make([]string, 0),
		RequiredSelection: make([]map[string]int, 0),
		CheckOption:       make([]views.KindOption, 0),
	},
}

const plus = 10

func init() {
	arr := MenuData[0].Items
	for i := 0; i < len(arr); i++ {
		arr[i].Prices[2] = arr[i].Prices[1] + plus
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
	return `<img src="/src/img/rice.jpg" alt="Smiley face">`
}
