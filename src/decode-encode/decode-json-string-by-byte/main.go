package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type contact struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Contact struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact"`
}

var jsonData = `{
	"name": "mrdulin",
	"title": "programmer",
	"contact": {
		"home": "shanghai",
		"cell": "222"
	}
}`

func main() {
	var contact contact
	err := json.Unmarshal([]byte(jsonData), &contact)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	contactFormatted, err := json.MarshalIndent(contact, "", "  ")
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	fmt.Println(string(contactFormatted))
}
