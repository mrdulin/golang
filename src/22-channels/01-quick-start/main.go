package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(1 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}

func main() {
	// var a chan int
	// if a == nil {
	// 	fmt.Println("channel a is nil, going to define it")
	// 	a = make(chan int)
	// 	fmt.Printf("Type of a is %T\n", a)
	// }

	done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	go hello(done)
	<-done
	fmt.Println("Main received data")

}
