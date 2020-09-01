package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	exitSignal := make(chan interface{})

	// create server to run on port the 9000
	server := &http.Server{
		Addr:    ":9000",
		Handler: nil, // use `DefaultServeMux`
	}

	// close server after 3 seconds
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("Close(): completed!", server.Close()) // close `server`
		close(exitSignal)                                  // close `exitSignal` channel
	})

	// launch server
	err := server.ListenAndServe()
	fmt.Println("ListenAndServe():", err)

	// listen to `exitSignal` channel
	<-exitSignal
	fmt.Println("Main(): exit complete!")
}
