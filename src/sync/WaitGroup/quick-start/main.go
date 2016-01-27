package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started goroutine: ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("goroutine %d ended\n", i)
	wg.Done()
}

func main() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		// 注意这里要传指针，不能传递值，否则每个goroutine将会有一个WaitGroup的拷贝，导致WaitGroup的counter不正确
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
