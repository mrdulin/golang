package main

import (
	"fmt"
	"sync"
)

var x int

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println("final value of x: ", x)
}

//final value of x:  937
//final value of x:  928
//final value of x:  923
//final value of x:  934
