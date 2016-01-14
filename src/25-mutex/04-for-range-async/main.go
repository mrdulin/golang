package main

import (
	"fmt"
	"sync"
	"time"
)

func updateStatus(id int, wg *sync.WaitGroup) {
	fmt.Printf("update status for id = %d\n", id)
	time.Sleep(time.Second * 3)
	defer wg.Done()
}

func updateAllStatus() error {
	ids := []int{1, 2, 3, 4}
	wg := sync.WaitGroup{}
	for _, id := range ids {
		wg.Add(1)
		go updateStatus(id, &wg)
	}
	wg.Wait()
	return nil
}

func main() {
	if err := updateAllStatus(); err != nil {
		fmt.Printf("update all status error = %v\n", err)
	}
	fmt.Printf("update all status done.")
}
