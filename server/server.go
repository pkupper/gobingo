package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func messageLoop(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error on reading message from client %s: %s", conn.RemoteAddr().String(), err.Error())
			return
		}

		if messageType != websocket.TextMessage {
			log.Printf("Unsupported message type from client %s", conn.RemoteAddr().String())
			return
		}

		handleMessage(p, conn)

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Printf("Error on sending response to client %s: %s", conn.RemoteAddr().String(), err.Error())
			return
		}
	}
}

func connectHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Client Connected: %s", conn.RemoteAddr().String())

	messageLoop(conn)
}

func main() {
	log.Println("Starting server...")

	http.HandleFunc("/connect", connectHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
