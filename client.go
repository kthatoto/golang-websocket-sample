package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	conn, _, _ := websocket.DefaultDialer.Dial("ws://localhost:8080/websocket", nil)
	defer conn.Close()

	go readMessage(conn)

	// Send message inputted from stdin
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		conn.WriteMessage(websocket.TextMessage, []byte(stdin.Text()))
		fmt.Printf("\x1b[34m     message wrote: %s \x1b[0m\n", stdin.Text())
	}
}

func readMessage(conn *websocket.Conn) {
	for {
		_, message, _ := conn.ReadMessage()
		fmt.Printf("\x1b[32m  message received: %s \x1b[0m\n", message)
	}
}
