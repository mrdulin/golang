package main

import (
	"fmt"
	"sync"
)

var x int

func increment(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<-ch
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var ch = make(chan bool, 1) // buffered channel with 1 capacity
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg, ch)
	}
	wg.Wait()
	fmt.Println("final value of x:", x) // final value of x: 1000
}
