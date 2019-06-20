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
	<li><a onclick="toManageShop()">Manage Menu</a>
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

	updateView.HTML = append(updateView.HTML, &resp.KeyValue{
		Key:  "Body",
		Data: "<h2>商店 " + shopName + "</h2>",
	})

	updateView.Script = append(updateView.Script, &resp.KeyValue{
		Key:  "src/js/manager/manageMenuMain.js",
		Data: "src/js/manager/manageMenuMain.js",
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
