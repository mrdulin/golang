package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	person := make(map[string]interface{})
	person["name"] = "mrdulin"
	person["title"] = "programmer"
	person["contact"] = map[string]interface{}{
		"home": "shanghai",
		"cell": "222",
	}

	data, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	fmt.Println(string(data))
}
