package main

import (
	"fmt"
	"sync"
)

var x int

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		// mutex必须传递指针，如果传递值，每个increment方法将会得到一份mutex的拷贝，goroutine的竞争问题依旧会出现
		go increment(&wg, &m)
	}
	wg.Wait()
	fmt.Println("final value of x: ", x) //final value of x:  1000
}
