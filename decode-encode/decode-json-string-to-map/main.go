package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var jsonData = `{
	"name": "mrdulin",
	"title": "programmer",
	"contact": {
		"home": "shanghai",
		"cell": "222"
	}
}`

func main() {
	var person map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	fmt.Println("Name:", person["name"])
	fmt.Println("Title:", person["title"])
	fmt.Println("Contact")
	fmt.Println("H:", person["contact"].(map[string]interface{})["home"])
	fmt.Println("C:", person["contact"].(map[string]interface{})["cell"])
}
