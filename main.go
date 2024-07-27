package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/streadway/amqp"
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

var mu sync.Mutex

var managerCount int
var volunteerCount int

type Message struct {
	MessageType int    `json:"messageType"`
	Body        string `json:"body"`
}

var ch *amqp.Channel

func handlevolunteerConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()

	mu.Lock()
	clients[ws] = true
	volunteerCount++
	fmt.Printf("New volunteer connection established. Total volunteer connections: %d\n", volunteerCount)
	mu.Unlock()

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			mu.Lock()
			volunteerCount--
			fmt.Printf("Volunteer connection Disconnected. Total volunteer connections: %d\n", volunteerCount)
			delete(clients, ws)
			mu.Unlock()
			break
		}
	}
}

func handleMessages() {
	if ch == nil {
		log.Fatal("RabbitMQ channel is not initialized")
	}

	q, err := ch.QueueDeclare(
		"messages",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	for d := range msgs {
		msg := Message{Body: string(d.Body)}
		mu.Lock()
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				fmt.Println(err)
				client.Close()
				volunteerCount--
				fmt.Printf("Volunteer connection closed. Total volunteer connections: %d\n", volunteerCount)
				delete(clients, client)
			}
		}
		mu.Unlock()
	}
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
	managerCount++
	fmt.Printf("New manager connection established. Total manager connections: %d\n", managerCount)
	mu.Unlock()

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			mu.Lock()
			delete(managers, ws)
			managerCount--
			fmt.Printf("Manager connection Disconnected. Total manager connections: %d\n", managerCount)
			mu.Unlock()
			break
		}

		err = ch.Publish(
			"",
			"messages",
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(msg.Body),
			})
		if err != nil {
			fmt.Println("Failed to publish a message:", err)
		}
	}
}

func loginasManager(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login as Manager")
}

func loginasvolunteer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login as Volunteer")
}

func getVolunteersCount(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count := volunteerCount
	mu.Unlock()

	fmt.Fprintf(w, "Total Volunteers: %d", count)
}

func getManagersCount(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count := managerCount
	mu.Unlock()

	fmt.Fprintf(w, "Total Managers: %d", count)
}

func setuproutes() {
	http.HandleFunc("/manager", handleManagerConnections)
	http.HandleFunc("/volunteer", handlevolunteerConnections)
	http.HandleFunc("/loginasmanager", loginasManager)
	http.HandleFunc("/loginasvolunteer", loginasvolunteer)
	http.HandleFunc("/volunteers/count", getVolunteersCount) // New endpoint for volunteers count
	http.HandleFunc("/managers/count", getManagersCount)     // New endpoint for managers count
}

func main() {
	fmt.Println("Go Websockets")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err = conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	setuproutes()
	go handleMessages()
	err = http.ListenAndServe(":8070", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
