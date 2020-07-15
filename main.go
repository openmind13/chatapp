package main

import (
	"fmt"
	"net/http"

	"github.com/openmind13/chatapp/handlers"
	"github.com/openmind13/chatapp/ws"
)

var (
	address = ":8080"
)

func main() {
	// fmt.Println("start chat app")

	mux := http.NewServeMux()

	staticFiles := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFiles))

	// url handlers
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/chat", handlers.ChatHandler)
	mux.HandleFunc("/signup", handlers.SignupHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)

	// Handle websocket connection
	mux.HandleFunc("/ws", ws.WebsocketListener)

	server := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	fmt.Printf("starting server: %s\n\n", server.Addr)

	server.ListenAndServe()
}
