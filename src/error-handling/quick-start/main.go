package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fileContent, err := loadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileContent)
}

func loadFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	return string(content), nil
}
