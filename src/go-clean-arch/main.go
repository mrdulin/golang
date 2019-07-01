package main

import (
	"fmt"
	cloudrun "go-clean-arch/infrastructure/cloudrun/reports"
	"log"
	"net/http"
	"os"
)

var port string

func init() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
}

func main() {
	http.HandleFunc("/google/reports/ad", cloudrun.GetAdPerformanceReport)

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
