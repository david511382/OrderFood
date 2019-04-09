package ws

import (
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/gin-gonic/gin"
)

var clients []*websocket.Conn = make([]*websocket.Conn, 0)
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Connect(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	clients = append(clients, ws)
}

func Notify(msg string) {
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
