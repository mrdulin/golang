package main

import (
	"fmt"
	"time"
)

func doSomething() {
	fmt.Println("do something")
}

func main() {
	go doSomething()

	time.Sleep(1 * time.Second)
	fmt.Println("main function")
}
