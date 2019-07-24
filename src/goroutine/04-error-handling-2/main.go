package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

func updateStatusTransaction(id int) (*int, error) {
	log.Printf("update status for id = %d", id)
	time.Sleep(time.Second * 5)
	if id > 2 {
		return nil, errors.New(fmt.Sprintf("error happened with id = %d", id))
	}
	return &id, nil
}

func updateAllStatus() error {
	ids := []int{1, 2, 3, 4, 5, 6}
	wg := sync.WaitGroup{}
	wg.Add(len(ids))
	for _, id := range ids {
		// 注意：要将for...range迭代的值传入go func() {}()闭包中，本例传入id
		go func(id int) {
			defer wg.Done()
			_, err := updateStatusTransaction(id)
			if err != nil {
				log.Printf("update status transaction error for id = %d. error = %v", id, err)
			}
		}(id)
	}
	wg.Wait()
	return nil
}

func main() {
	log.Printf("Start")
	err := updateAllStatus()
	if err != nil {
		log.Printf("update all status error = %v", err)
	}
	log.Print("End")
}
