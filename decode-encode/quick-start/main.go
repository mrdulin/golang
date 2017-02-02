package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type book struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Isbn13   string `json:"isbn13"`
	Price    string `json:"price"`
	Image    string `json:"image"`
	URL      string `json:"url"`
}

type searchResponse struct {
	Error string `json:"error"`
	Total string `json:"total"`
	Page  string `json:"page"`
	Books []book `json:"books"`
}

func main() {
	uri := `https://api.itbook.store/1.0/search/mongodb`

	resp, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	defer resp.Body.Close()

	var sr *searchResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&sr)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	srIndent, err := json.MarshalIndent(sr, "", "  ")

	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	fmt.Println(string(srIndent))
}
