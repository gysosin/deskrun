package main

import (
	"fmt"
	"log"
	"net/http"

	ws "github.com/gysosin/deskrun/backend/internal/websocket"
)

func main() {
	http.HandleFunc("/ws", ws.HandleConnection)

	port := 8080
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Server listening on http://localhost%s/ws", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
