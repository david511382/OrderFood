package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MenuItem struct {
	Name   string
	Prices []int
}

type KindOption struct {
	Name  string
	Price int
}

type MenuKind struct {
	Items             []MenuItem
	Size              []string
	RequiredSelection []map[string]int
	CheckOption       []KindOption
}

func Order(c *gin.Context, contains ...func() string) {
	c.Writer.WriteHeader(http.StatusOK)

	head := headHTML()
	c.Writer.Write([]byte(head))

	for _, f := range contains {
		c.Writer.Write([]byte(f()))
	}

	tail := tailHTML()
	c.Writer.Write([]byte(tail))
}

func headHTML() string {
	return `
		<html>
			<head>
				<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
				<link rel="stylesheet" type="text/css" href="src/css/style.css">
			</head>
			<body>
				<h1>Order Food</h1>`
}

func NameHTML() string {
	return `</br>使用者:<font id="name"></font></br>
	<script>
		$.ajax({
			type:"POST",
			url: "/get/name"
		}).done(function(name){
			$("#name").text(name);
		});    
	</script>
`
}

func FormHeadHTML() string {
	return `<form id="orderForm" name="orders">`
}

func MenuHTML() string {
	return `<hr><div id="menu"></div><hr>`
}

func ShopcartHTML() string {
	return `
		</br>
		購物清單:
		<div id="shopcart"></div>
		</br>
		總計:<font id="total"></font> 元
		</br>
		<script src="/src/js/menuView.js"></script>`
}

func FormEndHTML() string {
	return `<button id="SubmitButton"  type="button">送出</button>
	</form>
	<script src="/src/js/post.js"></script>
	`
}

func ResultHTML() string {
	return `<hr>
	全部訂單</br>
<textarea id="result" readonly></textarea>
`
}

func tailHTML() string {
	return `
		</body>	
		<script src="/src/js/menu.js"></script>
		<script src="/src/js/websocket.js"></script>
		</html>`
}
