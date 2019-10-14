package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create(".")
	if err != nil {
		panic("cannot create file")
	}
	defer file.Close()

	fmt.Fprintf(file, "hello")
}
