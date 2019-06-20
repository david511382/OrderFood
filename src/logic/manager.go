package logic

import (
	"fmt"
	"orderfood/src/database"
	"orderfood/src/database/models"
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
		%s
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

	js = `
	<script>
	function toHome(){
		$.ajax({
			type:"GET",
			url: "/manager"
		}).done(changePage);
	}

	function toManageShop(shopID){
		var url =  "/manager/managemenu?shopID=";
		if (shopID !== undefined) {
			url += shopID;
		}
		
		$.ajax({
			type:"GET",
			url: url
		}).done(changePage);
	}

	function changePage(html){
		document.body.innerHTML = html;
	}
	</script>
	`

	return
}

func ManageMenuView(shopID int32) (string, error) {
	html := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>%s</title>

        <link rel="stylesheet" type="text/css" href="/css/managerHome.css">
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
			<h2>商店 %s</h2>
		</div>

		<script>
			var shopData;
			var selectedShopID = %d;

			$.ajax({
				type:"GET",
				url: "/menu/shopmenu"
			}).done(init);

			function init(data){
				shopData = data
			}
		</script>
		%s
    </body>
    </html>
    `

	db := database.Db.MenuShop()
	shops, err := db.GetShop(nil)
	if err != nil {
		return "", err
	}

	treeHTML, treeJS := menuTree(shops)

	shopName := ""
	for _, shop := range shops {
		if shopID == 0 || shop.GetID() == shopID {
			shopName = shop.GetName()
			break
		}
	}

	html = fmt.Sprintf(html,
		"OrderFood後台",
		"Manage Menu",
		treeHTML,
		shopName,
		shopID,
		treeJS,
	)
	return html, nil
}
