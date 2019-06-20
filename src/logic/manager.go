package logic

import (
	"fmt"
	"orderfood/src/database"
	"orderfood/src/database/models"
	"strconv"
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

        <link rel="stylesheet" type="text/css" href="/css/managerHome.css">
        <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
    </head>
    <body>
        <h1>%s</h1>
		
		%s

        <form id="selectViewForm">
           %s
        </form>

        users order</br>
        <textarea class="list" id="userOrders" readonly></textarea>
        </br>

        total</br>
        <textarea class="list" id="result" readonly></textarea>

        <script src="/src/js/post.js"></script>
        <script src="/src/js/websocket.js"></script>
        <script src="/src/js/manager.js"></script>
    </body>
    </html>
    `

	db := database.Db.Menu()
	shops, err := db.GetShop(nil)
	if err != nil {
		return "", err
	}

	viewSelect := viewSelect(shops)

	tree := menuTree(shops)

	html = fmt.Sprintf(html,
		"OrderFood後台",
		"後台",
		viewSelect,
		tree,
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

func menuTree(shops []*models.Shop) string {
	shopStr := ""
	for _, shop := range shops {
		shopStr += `<li onclick="toManageShop(` + strconv.Itoa(int(shop.GetID())) + `)">` + shop.GetName() + "</li>"
	}

	result := `
	<ul id="myUL">
	<li onclick="toHome()">Home</li>
	<li><a onclick="toManageShop()">Manage Menu</a>
	  <ul>
		%s
	  </ul>
	</li>
	</ul>

	<script>
		function toHome(){
			$.ajax({
				type:"GET",
				url: "/manager"
			}).done(changePage);
		}

		function toManageShop(o){
			var url =  "/manager/managemenu?shopID=";
			if (o !== undefined) {
				url += o.innerHTML;
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

	result = fmt.Sprintf(result,
		shopStr,
	)
	return result
}

func ManageMenuView(shopID int) (string, error) {
	html := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>%s</title>

        <link rel="stylesheet" type="text/css" href="/css/managerHome.css">
        <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
    </head>
    <body>
        <h2>商店 %s</h2>
		%s

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
    </body>
    </html>
    `

	db := database.Db.Menu()
	shops, err := db.GetShop(nil)
	if err != nil {
		return "", err
	}

	tree := menuTree(shops)

	html = fmt.Sprintf(html,
		"OrderFood後台",
		"test",
		tree,
		shopID,
	)
	return html, nil
}
