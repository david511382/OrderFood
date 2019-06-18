package logic

import (
	"fmt"
	"orderfood/src/database"
	"strconv"
	"strings"
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
            <select id="viewSelect" name="view">
                <option value ="vag">素食</option>
                <option value ="rice">飯捲</option>
            </select>
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

	tree, err := menuTree()
	if err != nil {
		return "", err
	}

	html = fmt.Sprintf(html,
		"OrderFood後台",
		"後台",
		tree,
	)
	return html, nil
}

func menuTree() (string, error) {
	db := database.Db.Menu()
	shops, err := db.GetShop(nil)
	if err != nil {
		return "", err
	}

	shopArr := make([]string, 0)
	for _, shop := range shops {
		s := `<li onclick="toManageShop(` + strconv.Itoa(shop.GetID()) + `)">` + shop.GetName()
		shopArr = append(shopArr, s)
	}
	shopStr := strings.Join(shopArr, "</li>")
	shopStr = shopStr + "</li>"

	result := `
	<ul id="myUL">
	<li onclick="toHome()">Home</li>
	<li><a onclick="toManageShop()">Manage Shop</a>
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
			var url =  "/manager/manageshop?shopID=";
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
	return result, nil
}

func ManageShopView(shopID int) (string, error) {
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

	tree, err := menuTree()
	if err != nil {
		return "", err
	}

	html = fmt.Sprintf(html,
		"OrderFood後台",
		"test",
		tree,
		shopID,
	)
	return html, nil
}
