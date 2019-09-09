package handlers

import (
	"encoding/json"
	"fmt"
	"log"

	"../clients"
	"github.com/gorilla/websocket"
)

type loginJSON struct {
	Username string
	Password string
}

// Login Handler
func Login(data json.RawMessage, client clients.Client) {
	var loginData loginJSON
	json.Unmarshal(data, &loginData)
	log.Printf("Login: %s %s", loginData.Username, loginData.Password)
	clients.ConnectedClients.Range(func(key, value interface{}) bool {
		value.(clients.Client).Conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Client logged in: %s %s", loginData.Username, loginData.Password)))
		return true
	})
}
