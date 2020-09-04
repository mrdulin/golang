package main

import (
	"fmt"
	"time"

	"github.com/mrdulin/golang/src/goroutine/08-google-search/google"
)

func main() {
	start := time.Now()
	//results, err := google.Search("golang")
	//results, err := google.SearchParallel("golang")
	results, err := google.SearchTimeout("golang", 80*time.Millisecond)
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed, err)
}
