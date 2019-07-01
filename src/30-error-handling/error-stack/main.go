package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

func main() {
	filename := "test.txt"
	fileContent, err := loadFile(filename)
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	fmt.Println(fileContent)
}

func loadFile(filename string) (fileContent string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", errors.Wrap(err, "open file error")
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", errors.WithMessage(err, "read file content error")
	}
	return string(content), nil
}
