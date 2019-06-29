package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("received a request")
	target := os.Getenv("TARGET")
	if target == "" {
		target = "world"
	}
	fmt.Fprintf(w, "Hello %s\n", target)

}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
