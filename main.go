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
const t = r

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
	
	router.StaticFile("src/js/menu.js", "src/js/menu.js")
	router.StaticFile("src/js/vag.js", "src/js/vag.js")
	router.StaticFile("src/js/rice.js", "src/js/rice.js")
	router.StaticFile("src/js/post.js", "src/js/post.js")
	router.StaticFile("src/js/websocket.js", "src/js/websocket.js")

	admin := router.Group("/")
	switch t {
	case r:
		admin.GET("", rice.View)
		router.POST("/get/menu", func(c *gin.Context) {
			c.JSON(http.StatusOK, rice.MenuData)
		})
	case v:
		admin.GET("", vag.View)
		router.POST("/get/menu", func(c *gin.Context) {
			c.JSON(http.StatusOK, vag.MenuData)
		})
	}

	router.POST("/post/order", order)
	router.POST("/get/name", getUserName)
	router.POST("/get/order", getTotalOrders)

	router.GET("/ws", func(c *gin.Context) {
		ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}

		clients = append(clients, ws)
	})

	router.Run(domain)
}

func test(c *gin.Context) {
	c.HTML(http.StatusOK, "test.html", gin.H{
		"title": "訂飯捲",
	})
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
		name = "宜廷"
	case "192.168.0.108":
		name = "雨叡"
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
	default:
		name = ip
	}

	return
}

var userOrders = make(map[string][]string)

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
