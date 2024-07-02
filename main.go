package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var managers = make(map[*websocket.Conn]bool)
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)
var mu sync.Mutex

type Message struct {
	MessageType int    `json:"messageType"`
	Body        string `json:"body"`
}

func handleManagerConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()

	mu.Lock()
	managers[ws] = true
	mu.Unlock()

	for {
		var msg Message

		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			mu.Lock()
			delete(managers, ws)
			mu.Unlock()
			break
		}
		broadcast <- msg
	}
}

func handleClientConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()

	mu.Lock()
	clients[ws] = true
	mu.Unlock()

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			mu.Lock()
			delete(clients, ws)
			mu.Unlock()
			break
		}
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		mu.Lock()
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				fmt.Println(err)
				client.Close()
				delete(clients, client)
			}
		}
		mu.Unlock()
	}
}

func setuproutes() {
	http.HandleFunc("/manager", handleManagerConnections)
	http.HandleFunc("/client", handleClientConnections)
}

func main() {
	fmt.Println("Go Websockets")
	setuproutes()
	go handleMessages()
	err := http.ListenAndServe(":8070", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
