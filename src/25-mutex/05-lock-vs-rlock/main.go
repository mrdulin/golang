package main

import (
	"fmt"
	"sync"
	"time"
)

// https://blog.golang.org/maps
// https://tour.golang.org/concurrency/9
func main() {
	a := 0
	lock := sync.RWMutex{}

	// Lock只允许一个goroutine写
	for i := 1; i < 10; i++ {
		go func(i int) {
			lock.Lock()
			fmt.Printf("Lock: from goroutine %d: a = %d\n", i, a)
			time.Sleep(time.Second)
			lock.Unlock()
		}(i)
	}

	// RLock允许多个goroutine同时读
	b := 0
	for i := 11; i < 20; i++ {
		go func(i int) {
			lock.RLock()
			fmt.Printf("RLock: from goroutine %d: b = %d\n", i, b)
			time.Sleep(time.Second)
			lock.RUnlock()
		}(i)
	}

	<-time.After(10 * time.Second)

}
