package main

import (
	"fmt"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Websocket Endpoint")
}

func setuproutes() {

	http.HandleFunc("/", homepage)
	http.HandleFunc("/ws", wsEndpoint)

}
func main() {

	/*
		conn, ch, q := ConnectRabbitMQ()
		defer conn.Close()
		defer ch.Close()

		// Publish a message

		for i := 0; i < 10; i++ {
			go publishMessage(ch, q, fmt.Sprintf("Hello World %d", i))
		}

		go ConsumeMessages(ch, q, 1)

		// Periodically check the status of the queue
		go func() {
			for {
				CheckQueueStatus(ch, q.Name)
				time.Sleep(10 * time.Second)
			}
		}()

		// Printing a message to the console and waiting for a signal to exit
		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		select {}
	*/

	fmt.Println("Go Websockets")
	setuproutes()
	http.ListenAndServe(":8080", nil)

}
