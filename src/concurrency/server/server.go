package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const Port string = "9090"

func CreateHTTPServer() error {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		values := req.URL.Query()
		echo := values["echo"][0]
		log.Printf("echo: %s", echo)
		time.Sleep(2 * time.Second)
		if _, err := fmt.Fprint(w, echo); err != nil {
			log.Fatal(err)
		}
	})

	log.Printf("Http server is listening on http://localhost:%s\n", Port)
	err := http.ListenAndServe(":"+Port, nil)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := CreateHTTPServer()
	if err != nil {
		log.Fatal(err)
	}
}
