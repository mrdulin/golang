package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/icrowley/fake"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("received a request")
	_, err := fmt.Fprintf(w, "email: %s\n", fake.EmailAddress())
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
