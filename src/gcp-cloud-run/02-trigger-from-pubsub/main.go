package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type PubSubMessage struct {
	Message struct {
		Data []byte `json:"data,omitempty"`
		ID   string `json:"id"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

func HelloPubSub(w http.ResponseWriter, r *http.Request) {
	var m PubSubMessage
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		log.Printf("json.NewDecoder: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	name := string(m.Message.Data)
	if name == "" {
		name = "World"
	}
	log.Printf("Hello %s!", name)
}

func main() {
	http.HandleFunc("/", HelloPubSub)
	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	// Start HTTP server.
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
