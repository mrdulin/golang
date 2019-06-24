package jsonextend

import (
	"encoding/json"
	"strings"
)

// DecodeJSONStream is used to decode json stream
func DecodeJSONStream(jsonStream string, message interface{}) ([]interface{}, error) {
	decoder := json.NewDecoder(strings.NewReader(jsonStream))

	_, err := decoder.Token()
	if err != nil {
		return nil, err
	}
	var messages []interface{}
	for decoder.More() {
		// var message Message
		err := decoder.Decode(&message)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	_, err = decoder.Token()
	if err != nil {
		return nil, err
	}
	return messages, nil
}