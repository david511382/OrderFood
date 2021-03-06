package manager

import (
	"fmt"
	"orderfood/src/database"
	"orderfood/src/database/models"
	"orderfood/src/handler/models/resp"
	"orderfood/src/logic"
	"strconv"

	linq "github.com/ahmetb/go-linq"
)

func ManagerView(username string) (string, error) {
	if username != "localhost" {
		return username + " 禁止進入!!", logic.DenyError
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

		<script src="/src/js/manager/api.js"></script>
		<script src="/src/js/ajax.js"></script>
        <script src="/src/js/post.js"></script>
        <script src="/src/js/websocket.js"></script>
		<script id="homeJS" src="/src/js/manager/home/main.js"></script>
		<script id="%s" src="%s"></script>
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
		treeJS.Key,
		treeJS.Data,
	)
	return html, nil
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
		Key:  treeJS.Key,
		Data: treeJS.Data,
	})

	return updateView, nil
}

func ManageMenuView(shopID int32) (*resp.UpdateView, error) {
	updateView := newUpdateView()

	updateView.Data = strconv.Itoa(int(shopID))

	updateView.HTML = append(updateView.HTML, &resp.KeyValue{
		Key:  "Header",
		Data: "<h1>Manage Menu</h1>",
	})

	bodyHTML := `
		<div id="menuHead">
			</br>	
			<a>商店</a>
			<input id="shopNameInput" type="text" onkeypress="shopNameInputKeyPress(this.value)" onfocus="this.select()"></input>			
			<button onclick="removeShopButtonClick()">刪除</button>
			</br>
			</br>

			<table id="optionTable" border="0">
			</table>
			</br>
		</div>

		<table border="0">
			<tr>
				<td>
					<a>選單</a>
					<a id="currentOptionNameA"></a>
					</br></br>
				</td>
				<td id="optionSelectTd"></td>
			</tr>	
			<tr>
				<td>
					<a>商品</a>
					<table id="itemTable" border="1">
						<tr>
							<td>所屬選單</td>
							<td>品名</td>
							<td>價格</td>
							<td>操作</td>
						</tr>
						<tr>
							<td id="newItemOptionNameTd">小,大</td>
							<td><input id="newItemNameInput"></input></td>
							<td><input id="newItemPriceInput"></input></td>
							<td><button id="addItemButton" onclick="newItemButtonClick()">新增</button></td>
						</tr>
						<tr>
							<td></td>
							<td><select id="newItemNameSelect"></select></td>
							<td></td>
							<td><button id="addItemButton">加入</button></td>
						</tr>
					</table>
				</td>
				<td id="selectionTableTd">
				</td>
			</tr>	
		</table>

		<div id="forOptionDiv"></div>
		`
	updateView.HTML = append(updateView.HTML, &resp.KeyValue{
		Key:  "Body",
		Data: bodyHTML,
	})

	updateView.Script = append(updateView.Script, &resp.KeyValue{
		Key:  "manageMenuJS",
		Data: "src/js/manager/manageMenu/main.js",
	})

	return updateView, nil
}

func NewShopView() (*resp.UpdateView, error) {
	updateView := newUpdateView()

	updateView.HTML = append(updateView.HTML, &resp.KeyValue{
		Key:  "Header",
		Data: "<h1>New Shop</h1>",
	})

	bodyHTML := `
		</br>	
		<a>商店</a>
		<input id="shopNameInput" type="text"></input>
		<button onclick="NewShopButtonClick()">新增</button>
		`
	updateView.HTML = append(updateView.HTML, &resp.KeyValue{
		Key:  "Body",
		Data: bodyHTML,
	})

	updateView.Script = append(updateView.Script, &resp.KeyValue{
		Key:  "newshopJS",
		Data: "src/js/manager/newshop/main.js",
	})

	return updateView, nil
}

func NewOptionView() (*resp.UpdateView, error) {
	updateView := newUpdateView()

	updateView.HTML = append(updateView.HTML, &resp.KeyValue{
		Key:  "Header",
		Data: "<h1>New Option</h1>",
	})

	headHTML := `
		</br>	
		<a>商店</a>
		<a id="shopNameInput"></a>			
		</br>
		</br>
		`
	updateView.HTML = append(updateView.HTML, &resp.KeyValue{
		Key:  "menuHead",
		Data: headHTML,
	})

	bodyHTML := `
		<button onclick="doneButtonClick()">完成</button>
		<button onclick="cancelButtonClick()">取消</button>
		`
	updateView.HTML = append(updateView.HTML, &resp.KeyValue{
		Key:  "forOptionDiv",
		Data: bodyHTML,
	})

	updateView.Script = append(updateView.Script, &resp.KeyValue{
		Key:  "newoptionJS",
		Data: "src/js/manager/newoption/main.js",
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

func menuTree(shops []*models.Shop) (html string, js *resp.KeyValue) {
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

	js = &resp.KeyValue{
		Key:  "menuTreeJS",
		Data: "/src/js/manager/treenode.js",
	}

	return
}
