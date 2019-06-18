package logic

import (
	"fmt"
	"orderfood/src/database"
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


<script>
var toggler = document.getElementsByClassName("caret");
var i;

for (i = 0; i < toggler.length; i++) {
toggler[i].addEventListener("click", function() {
  this.parentElement.querySelector(".nested").classList.toggle("active");
  this.classList.toggle("caret-down");
});
}
</script>

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

	const ToManageShopLIStr = `<li onclick="toManageShop(this)">`

	shopArr := make([]string, 0)
	for _, shop := range shops {
		shopArr = append(shopArr, shop.GetName())
	}
	shopStr := strings.Join(shopArr, "</li>"+ToManageShopLIStr)
	shopStr = ToManageShopLIStr + shopStr + "</li>"

	result := `
	<ul id="myUL">
	<li onclick="toHome()">Home</li>
	<li><span class="caret">Manage Shop</span>
	  <ul class="nested">
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
			var url =  "/manager/manageshop";
			if (o !== undefined) {
				url += "/" + o.innerHTML;
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
