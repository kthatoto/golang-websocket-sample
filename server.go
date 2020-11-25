package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/websocket", serveWebsocket)
	http.ListenAndServe(":8080", nil)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var connections map[*websocket.Conn]bool

func serveWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)

	// Connect and Handle disconnect
	connections[conn] = true
	defer func() {
		if _, ok := connections[conn]; ok {
			delete(connections, conn)
		}
	}()

	// Read message and Send it to all connections
	for {
		messageType, message, _ := conn.ReadMessage()
		for c := range connections {
			c.WriteMessage(messageType, message)
		}
	}
}
