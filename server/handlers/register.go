package handlers

import (
	"encoding/json"
	"log"

	"../clients"
)

type registerJSON struct {
	Username string
	Password string
}

// Register Handler
func Register(data json.RawMessage, client clients.Client) {
	var registerData registerJSON
	json.Unmarshal(data, &registerData)
	log.Printf("Login: %s %s", registerData.Username, registerData.Password)
}
