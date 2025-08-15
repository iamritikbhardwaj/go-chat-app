package main

import (
	"html/template"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	upgrader  = websocket.Upgrader{}
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan string)
	mu        = sync.Mutex{}
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	templ, _ := template.ParseFiles("templates/index.html")
	templ.Execute(w, nil)
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer ws.Close()

	mu.Lock()
	clients[ws] = true
	mu.Unlock()

	for {
		var message string
		err := ws.ReadJSON(&message)
		if err != nil {
			log.Println(err)
			mu.Lock()
			delete(clients, ws)
			mu.Unlock()
			break
		}
		broadcast <- message
	}
}

func handleMessage() { // brodcasting to all the clients
	for {
		msg := <-broadcast
		mu.Lock()
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
		mu.Unlock()
	}
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", handleConnections)

	go handleMessage()

	log.Println("Server running on port 8008")
	http.ListenAndServe(":8008", nil)
}
