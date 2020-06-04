package main

import (
	"fmt"
	"time"
)

func simpleSample() {
	ch := make(chan string, 2)
	ch <- "golang"
	ch <- "serverless"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func write(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

func complexSample() {
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)
	}
}

func deadLock() {
	ch := make(chan string, 2)
	ch <- "golang"
	ch <- "typescript"
	ch <- "serverless"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func capacityAndLength() {
	ch := make(chan string, 3)
	ch <- "golang"
	ch <- "typescript"
	fmt.Println("capacity is: ", cap(ch))
	fmt.Println("length is: ", len(ch))
	fmt.Println("read value", <-ch)
	fmt.Println("new length is", len(ch))
}

func main() {
	// simpleSample()
	complexSample()
	// deadLock()
	// capacityAndLength()
}
