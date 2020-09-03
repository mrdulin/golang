package main

import (
	"encoding/json"
	"fmt"

	"github.com/mrdulin/golang/src/decode-encode/decode-json-stream/jsonextend"
)

// Message is used to store json data
type Message struct {
	Name, Text string
}

func main() {
	const jsonStream = `
		
			{"Name": "Ed", "Text": "Knock knock."},
			{"Name": "Sam", "Text": "Who's there?"},
			{"Name": "Ed", "Text": "Go fmt."},
			{"Name": "Sam", "Text": "Go fmt who?"},
			{"Name": "Ed", "Text": "Go fmt yourself!"}
		]
	`

	var message = Message{}

	messages, err := jsonextend.DecodeJSONStream(jsonStream, message)
	if err != nil {
		if serr, ok := err.(*json.SyntaxError); ok {
			fmt.Printf("%#v\n", serr)
		} else {
			fmt.Printf("%#v", err)
		}
		return
	}
	fmt.Printf("%#v\n", messages)
}
