package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
			<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
						
			<body>
				<div>Order Food</div>`
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

func FormEndHTML() string {
	return `<button  type="submit">送出</button>
	</form>
	<script src="/src/js/post.js"></script>
	`
}

func ResultHTML() string {
	return `<hr>
	全部訂單</br>
<textarea cols="30" rows="10" id="result" readonly></textarea>
`
}

func tailHTML() string {
	return `
		</body>	
		<script src="/src/js/menu.js"></script>
		<script src="/src/js/websocket.js"></script>
		</html>`
}