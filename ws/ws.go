package ws

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	clients       = make(map[*websocket.Conn]bool)
	broadcastChan = make(chan Message)
)

// WebsocketListener ...
func WebsocketListener(w http.ResponseWriter, r *http.Request) {
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("could not upgrade the upgrader")
		fmt.Println(err)
		return
	}
	defer wsConn.Close()

	clients[wsConn] = true

	go serverSender(wsConn)

	go sendBroadcast()

	// listen messages from user

	for {
		var msg Message
		err := wsConn.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v\n", err)
			delete(clients, wsConn)
			break
		}

		fmt.Printf("%v -> %v\n", msg.Username, msg.Text)

		broadcastChan <- msg
	}
}

// take message from broadcast channel and send to all connections
func sendBroadcast() {
	for {
		msg := <-broadcastChan

		// erase uuid for security
		msg.UUID = ""

		for client := range clients {
			err := client.WriteJSON(&msg)
			if err != nil {
				log.Printf("error: %v\n", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

// websocketSender send message from server terminal
func serverSender(wsConn *websocket.Conn) {
	reader := bufio.NewReader(os.Stdin)

	for {
		rawText, _ := reader.ReadString('\n')
		text := strings.Replace(rawText, "\n", "", -1)

		msg := Message{
			Text:     text,
			Username: "server",
			UUID:     "",
		}

		// wsConn.WriteJSON(&msg)
		broadcastChan <- msg
	}
}
