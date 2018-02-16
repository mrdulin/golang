package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	file, err := os.Create(".")
	if err != nil {
		panic("cannot create file")
	}

	defer file.Close()

	fmt.Fprintf(file, "hello")

}
