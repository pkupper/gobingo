package main

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"

	"./clients"
	"./handlers"
)

type messageJSON struct {
	Type    messageType
	Message json.RawMessage
}

type messageType int32

const (
	register messageType = 0
	login    messageType = 1
)

func handleMessage(message []byte, conn *websocket.Conn) {
	client, ok := clients.ConnectedClients.Load(conn.RemoteAddr().String())
	if !ok {
		client = clients.Client{Conn: conn, Name: ""}
		clients.ConnectedClients.Store(conn.RemoteAddr().String(), client)
	}
	var messageObject messageJSON

	json.Unmarshal(message, &messageObject)

	switch messageObject.Type {
	case register:
		handlers.Register(messageObject.Message, client.(clients.Client))
	case login:
		handlers.Login(messageObject.Message, client.(clients.Client))
	default:
		log.Println("Error, unknown message type")
	}
}
