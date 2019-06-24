package jsonextend

import "testing"

type decodeJSONStreamTest struct {
	jsonStream string
	expectedErrorMessage string
}

func TestDecodeJSONStream(t *testing.T) {
	var message = struct {
		Name, Text string
	}{}

	var decodeJSONStreamTests = []decodeJSONStreamTest{
		{
			jsonStream: `
				
					{"Name": "Ed", "Text": "Knock knock."},
					{"Name": "Sam", "Text": "Who's there?"},
					{"Name": "Ed", "Text": "Go fmt."},
					{"Name": "Sam", "Text": "Go fmt who?"},
					{"Name": "Ed", "Text": "Go fmt yourself!"}
				]
			`,
			expectedErrorMessage: "not at beginning of value",
		},
		{
			jsonStream: `
				[		
					{"Name": "Ed", "Text": "Knock knock."},
					{"Name: "Sam", "Text": "Who's there?"},
					{"Name": "Ed", "Text": "Go fmt."},
					{"Name": "Sam", "Text": "Go fmt who?"},
					{"Name": "Ed", "Text": "Go fmt yourself!"}
				]
			`,
			expectedErrorMessage: "invalid character 'S' after object key",
		},
	}

	for _, testCase := range decodeJSONStreamTests {
		_, err := DecodeJSONStream(testCase.jsonStream, message)	
		if err != nil {
			if err.Error() != testCase.expectedErrorMessage {
				t.Errorf("%#v", err)
			}
		}
	}
	
}