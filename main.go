package main

import (
	"fmt"
	"net/http"
	rice "orderfood/src/views/Rice"
	vag "orderfood/src/views/Vag"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"

	"github.com/gin-gonic/gin"
)

type serverConfig struct {
	Domain string `json:"Hostname"`
}

const domain = "192.168.0.144:5487"

//const domain = "localhost:5487"
const r = "rice"
const v = "vag"

var targetView = v

var userOrders = make(map[string][]string)

var clients []*websocket.Conn = make([]*websocket.Conn, 0)
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("src/templates/*.html")
	router.StaticFile("src/img/rice.jpg", "src/img/rice.jpg")
	router.StaticFile("src/img/vag.jpg", "src/img/vag.jpg")
	router.StaticFile("favicon.ico", "src/img/head.ico")

	router.StaticFile("src/css/style.css", "src/css/style.css")
	router.StaticFile("src/css/managerStyle.css", "src/css/managerStyle.css")

	router.StaticFile("src/js/menu.js", "src/js/menu.js")
	router.StaticFile("src/js/menuView.js", "src/js/menuView.js")
	router.StaticFile("src/js/post.js", "src/js/post.js")
	router.StaticFile("src/js/manager.js", "src/js/manager.js")
	router.StaticFile("src/js/websocket.js", "src/js/websocket.js")

	admin := router.Group("/")
	admin.GET("", index)

	router.GET("/manager", manager)

	router.POST("/get/menu", getMenu)
	router.POST("/post/order", order)
	router.POST("/get/name", getUserName)
	router.POST("/get/order", getTotalOrders)
	router.POST("/get/user/orders", userOrder)
	router.POST("/post/view", changeView)

	router.GET("/ws", func(c *gin.Context) {
		ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}

		clients = append(clients, ws)
	})

	router.Run(domain)
}

func index(c *gin.Context) {
	switch targetView {
	case r:
		rice.View(c)
	case v:
		vag.View(c)
	}
}

func manager(c *gin.Context) {
	user := mapUserName(c)
	if user != "哥" {
		c.String(http.StatusOK, user+"禁止進入")
		return
	}

	c.HTML(http.StatusOK, "manager.html", gin.H{
		"title": "後台",
	})
}

func getMenu(c *gin.Context) {
	switch targetView {
	case r:
		c.JSON(http.StatusOK, rice.MenuData)
	case v:
		c.JSON(http.StatusOK, vag.MenuData)
	}
}

func getUserName(c *gin.Context) {
	name := mapUserName(c)

	c.String(http.StatusOK, name)
}

func mapUserName(c *gin.Context) (name string) {
	switch ip := c.Request.RemoteAddr[:13]; ip {
	case "192.168.0.106":
		name = "奇異"
	case "192.168.0.121":
		name = "QQ詩"
	case "192.168.0.135":
		name = "廷"
	case "192.168.0.108":
		name = "叡"
	case "192.168.0.137":
		name = "雞排"
	case "192.168.0.115":
		name = "傑哥"
	case "192.168.0.144":
		name = "哥"
	case "192.168.0.122":
		name = "宏哥"
	case "192.168.0.104":
		name = "小麥"
	case "192.168.0.123":
		name = "hank"
	case "192.168.0.136":
		name = "Indy"
	case "192.168.0.128":
		name = "尾"
	default:
		name = ip
	}

	return
}

func changeView(c *gin.Context) {
	targetView = c.PostForm("view")
	c.String(http.StatusOK, targetView)
}

func userOrder(c *gin.Context) {
	result := ""
	for name, orders := range userOrders {
		result += name + "\n"
		for _, order := range orders {
			result += order + "\n"
		}
		result += "\n"
	}

	c.String(http.StatusOK, result)
}

func order(c *gin.Context) {
	orderStr := c.PostForm("orders")

	name := mapUserName(c)

	orders := strings.Split(orderStr, "\n")
	orders = orders[:len(orders)-1]
	userOrders[name] = orders

	log := fmt.Sprintf(
		"%s\n%s\n",
		name, orderStr,
	)
	fmt.Print(log)

	totalList := integrationOrders()
	c.String(http.StatusOK, totalList)

	notify(totalList)
}

func notify(msg string) {
	message := []byte(msg)

	for id, ws := range clients {
		w, err := ws.NextWriter(websocket.TextMessage)
		if err != nil {
			clients[id] = nil
			continue
		}

		if _, err := w.Write(message); err != nil {
			clients[id] = nil
			continue
		}

		if err := w.Close(); err != nil {
			clients[id] = nil
			continue
		}
	}

	existID := 0
	for j := 1; existID < len(clients); existID++ {
		if clients[existID] == nil {
			next := existID + j
			if next >= len(clients) {
				break
			}

			if clients[next] != nil {
				clients[existID] = clients[next]
				clients[next] = nil
			} else {
				j++
				existID--
			}
		} else {
			j = 1
		}
	}
	clients = clients[:existID]
}

func getTotalOrders(c *gin.Context) {
	totalList := integrationOrders()

	c.String(http.StatusOK, totalList)
}

func integrationOrders() (totalList string) {
	totalOrders := make(map[string]int)

	for _, orders := range userOrders {
		for _, order := range orders {
			orderElements := strings.Split(order, " ")
			amount, _ := strconv.Atoi(orderElements[len(orderElements)-3])
			orderElements = orderElements[:len(orderElements)-3]

			clearOrder := strings.Join(orderElements, " ")

			totalOrders[clearOrder] += amount
		}
	}

	for order, amount := range totalOrders {
		totalList += order + " * " + strconv.Itoa(amount) + "\n"
	}

	return
}
