package logic

import (
	"fmt"
	"orderfood/src/database"
	"orderfood/src/database/models"
	"orderfood/src/handler/models/resp"
	"strconv"

	linq "github.com/ahmetb/go-linq"
)

func ManagerView(username string) (string, error) {
	if username != "localhost" {
		return username + " 禁止進入!!", DenyError
	}

	html := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>%s</title>

        <link rel="stylesheet" type="text/css" href="/css/manager.css">
        <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
    </head>
	<body>
		<div id="Header">
        	<h1>%s</h1>
		</div>

		<div id="Sidebar">
		%s
		</div>

		<div id="Body">
			<form id="selectViewForm">
			%s
			</form>
		
			users order</br>
			<textarea class="list" id="userOrders" readonly></textarea>
			</br>

			total</br>
			<textarea class="list" id="result" readonly></textarea>
		</div>

        <script src="/src/js/post.js"></script>
        <script src="/src/js/websocket.js"></script>
		<script src="/src/js/manager.js"></script>
		<script src="/src/js/manager/api.js"></script>
		<script src="%s"></script>
    </body>
    </html>
    `

	db := database.Db.Menu()
	shops, err := db.GetShop(nil)
	if err != nil {
		return "", err
	}

	viewSelect := viewSelect(shops)

	treeHTML, treeJS := menuTree(shops)

	html = fmt.Sprintf(html,
		"OrderFood後台",
		"後台",
		treeHTML,
		viewSelect,
		treeJS,
	)
	return html, nil
}

func viewSelect(shops []*models.Shop) string {
	shopStr := ""
	for _, shop := range shops {
		shopStr += `<option value ="` + strconv.Itoa(int(shop.GetID())) + `">` + shop.GetName() + `</option>`
	}

	result := `
	<select id="viewSelect" name="view">
		%s
	</select>`

	result = fmt.Sprintf(result,
		shopStr,
	)
	return result
}

func MenuTreeView() (*resp.UpdateView, error) {
	db := database.Db.Menu()
	shops, err := db.GetShop(nil)
	if err != nil {
		return nil, err
	}

	treeHTML, treeJS := menuTree(shops)

	updateView := newUpdateView()

	updateView.HTML = append(updateView.HTML, &resp.KeyValue{
		Key:  "Sidebar",
		Data: treeHTML,
	})

	updateView.Script = append(updateView.Script, &resp.KeyValue{
		Key:  treeJS,
		Data: treeJS,
	})

	return updateView, nil
}

func menuTree(shops []*models.Shop) (html string, js string) {
	linq.From(shops).OrderBy(func(m interface{}) interface{} {
		shop := m.(*models.Shop)
		return shop.GetID()
	}).ToSlice(&shops)

	shopStr := ""
	for _, shop := range shops {
		shopStr += `<li onclick="toManageShop(` + strconv.Itoa(int(shop.GetID())) + `)">` + shop.GetName() + "</li>"
	}

	html = `
	<ul id="myUL">
	<li onclick="toHome()">Home</li>
	<li><a onclick="toNewShop()">New Shop</a>
	  <ul>
		%s
	  </ul>
	</li>
	</ul>
	`

	html = fmt.Sprintf(html,
		shopStr,
	)

	js = "/src/js/manager/treenode.js"

	return
}

func ManageMenuView(shopID int32) (*resp.UpdateView, error) {
	updateView := newUpdateView()

	updateView.HTML = append(updateView.HTML, &resp.KeyValue{
		Key:  "Header",
		Data: "<h1>Manage Menu</h1>",
	})

	db := database.Db.MenuShop()
	shops, err := db.GetShop(&models.Shop{ID: shopID})
	if err != nil {
		return nil, err
	}

	shopName := ""
	if len(shops) != 0 {
		shopName = shops[0].GetName()
	}

	bodyHTML := `
		</br>	
		<a>商店</a>
		<textarea id="shopNameTextarea">%s</textarea>
		<button id="addShopButton">+</button>
		<button id="removeShopButton">-</button>
		</br>
		</br>

		<a>選單</a>
		<a>小,大</a>
		</br>
		</br>

		<button id="allOptionButton">所有</button>
		<button id="noneOptionButton">無</button>
		<button id="addOptionButton">+</button>
		</br>
		</br>

		<a>商品</a>
		<table border="1">
			<tr>
				<td>所屬選單</td>
				<td>品名</td>
				<td>價格</td>
				<td>操作</td>
			</tr>
			<tr>
				<td>小,大|辣油</td>
				<td>炒麵</td>
				<td>15</td>
				<td><button>刪除</button></td>
			</tr>
			<tr>
				<td></td>
				<td><select id="newItemNameSelect"></select></td>
				<td></td>
				<td><button id="addItemButton">加入</button></td>
			</tr>
			<tr>
				<td>小,大</td>
				<td><textarea id="newItemNameTextarea"></textarea></td>
				<td><textarea id="newItemPriceTextarea"></textarea></td>
				<td><button id="addItemButton">新增</button></td>
			</tr>
		</table>
		</br>
		</br>

		<a>選單選項</a>
		<table border="1">
			<tr>
				<td>名稱</td>
				<td>加價</td>
				<td>操作</td>
			</tr>
			<tr>
				<td>小</td>
				<td>0</td>
				<td><button>刪除</button></td>
			</tr>
			<tr>
				<td>大</td>
				<td>5</td>
				<td><button>刪除</button></td>
			</tr>
			<tr>
				<td><textarea id="newSelectionNameTextarea"></textarea></td>
				<td><textarea id="newSelectionPriceTextarea"></textarea></td>
				<td><button id="addSelectionButton">新增</button></td>
			</tr>
		</table>
		`
	bodyHTML = fmt.Sprintf(bodyHTML, shopName)
	updateView.HTML = append(updateView.HTML, &resp.KeyValue{
		Key:  "Body",
		Data: bodyHTML,
	})

	updateView.Script = append(updateView.Script, &resp.KeyValue{
		Key:  "src/js/manager/manageMenuMain.js",
		Data: "src/js/manager/manageMenuMain.js",
	})

	return updateView, nil
}

func NewShopView() (*resp.UpdateView, error) {
	updateView := newUpdateView()

	bodyHTML := `
		</br>	
		<a>商店</a>
		<textarea id="shopNameTextarea"></textarea>
		<button onclick="NewShopButtonClick()">新增</button>
		`
	updateView.HTML = append(updateView.HTML, &resp.KeyValue{
		Key:  "Body",
		Data: bodyHTML,
	})

	updateView.Script = append(updateView.Script, &resp.KeyValue{
		Key:  "src/js/manager/newshop.js",
		Data: "src/js/manager/newshop.js",
	})

	return updateView, nil
}

func newUpdateView() *resp.UpdateView {
	return &resp.UpdateView{
		HTML:   make([]*resp.KeyValue, 0),
		Script: make([]*resp.KeyValue, 0),
		Css:    make([]*resp.KeyValue, 0),
	}
}
