package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// TODO: tighten this in production
		return true
	},
}

// HandleConnection upgrades the HTTP connection and processes WebSocket messages
func HandleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}
	defer conn.Close()

	log.Println("WebSocket client connected")

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("read error: %v", err)
			break
		}

		log.Printf("received: %s", msg)

		// Echo message for now
		err = conn.WriteMessage(msgType, msg)
		if err != nil {
			log.Printf("write error: %v", err)
			break
		}
	}
}
