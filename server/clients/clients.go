package clients

import (
	"sync"

	"github.com/gorilla/websocket"
)

// Client Struct
type Client struct {
	Conn *websocket.Conn
	Name string
}

var ConnectedClients sync.Map
