package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestWebSocketEcho(t *testing.T) {
	// Start test server
	server := httptest.NewServer(http.HandlerFunc(handleConnections))
	defer server.Close()

	// Upgrade HTTP to WebSocket URL
	u := "ws" + server.URL[len("http"):] + "/ws"

	// Connect WebSocket client
	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fatalf("Failed to connect to WebSocket: %v", err)
	}
	defer ws.Close()

	// Start broadcast listener
	// go handleMessages()

	// Send message
	message := "Hello, Test!"
	err = ws.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	// Read response (broadcast)
	ws.SetReadDeadline(time.Now().Add(2 * time.Second)) // avoid hanging
	_, resp, err := ws.ReadMessage()
	if err != nil {
		t.Fatalf("Failed to read message: %v", err)
	}

	if string(resp) != message {
		t.Errorf("Expected message %q, got %q", message, string(resp))
	}
}
