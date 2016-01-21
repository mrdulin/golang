package main

import (
	"fmt"
	"time"
)

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

func main() {
	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1 * time.Second)
		select {
		case rval := <-ch:
			fmt.Println(rval)
			return
		default:
			fmt.Println("no value received")
		}
	}
}
