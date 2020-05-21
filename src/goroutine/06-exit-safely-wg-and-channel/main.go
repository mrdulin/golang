package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int, ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		default:
			fmt.Printf("worker %d: hello\n", i)
		case <-ch:
			fmt.Printf("worker: %d exited\n", i)
			return
		}
	}
}

func main() {
	ch := make(chan bool)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}
	time.Sleep(time.Second)
	close(ch)
	wg.Wait()
}
