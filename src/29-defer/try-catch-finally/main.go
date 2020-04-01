package main

import (
	"fmt"
	"os"
)

func main() {
	// catch
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	// try
	file, err := os.Create(".")
	if err != nil {
		panic("cannot create file")
	}

	// finally
	defer file.Close()

	fmt.Fprintf(file, "hello")

}
