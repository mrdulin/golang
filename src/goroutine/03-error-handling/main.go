package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

type updateStatusResult struct {
	Payload interface{}
	Error   error
}

func updateStatus(id int, ch chan updateStatusResult, wg *sync.WaitGroup) {
	//fmt.Printf("update status with id = %d\n", id)
	defer wg.Done()
	if id > 2 {
		ch <- updateStatusResult{Payload: nil, Error: errors.New(fmt.Sprintf("update status with id = %d error", id))}
		return
	}
	ch <- updateStatusResult{Payload: fmt.Sprintf("update status with id = %d success", id), Error: nil}
}

func updateAllStatus() []updateStatusResult {
	ids := []int{1, 2, 3, 4}
	var results []updateStatusResult
	ch := make(chan updateStatusResult)
	wg := sync.WaitGroup{}
	wg.Add(len(ids))
	for _, id := range ids {
		go updateStatus(id, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		results = append(results, result)
	}

	return results
}

func main() {
	results := updateAllStatus()
	//log.Printf("%+v", results)
	for _, result := range results {
		if result.Error != nil {
			log.Printf("%v\n", result.Error)
			continue
		}
		log.Printf("%v\n", result.Payload)
	}
}
