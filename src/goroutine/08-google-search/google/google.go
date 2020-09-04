package google

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var (
	Web   = FakeSearch("web", "The Go Programming Language", "http://golang.org")
	Image = FakeSearch("image", "The Go gopher", "https://blog.golang.org/gopher/gopher.png")
	Video = FakeSearch("video", "Concurrency is not Parallelism", "https://www.youtube.com/watch?v=")
)

type Result struct {
	Title string
	URL   string
}

type SearchFunc func(query string) Result

func FakeSearch(kind, title, url string) SearchFunc {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result{
			Title: fmt.Sprintf("%s(%q): %s", kind, query, title),
			URL:   url,
		}
	}
}

func Search(query string) ([]Result, error) {
	results := []Result{
		Web(query),
		Image(query),
		Video(query),
	}
	return results, nil
}

func SearchParallel(query string) ([]Result, error) {
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()
	return []Result{<-c, <-c, <-c}, nil
}

func SearchTimeout(query string, timeout time.Duration) ([]Result, error) {
	timer := time.After(timeout)
	c := make(chan Result, 3)
	go func() {
		fmt.Println("go web")
		c <- Web(query)
		fmt.Println("go web end")
	}()
	go func() {
		fmt.Println("go image")
		c <- Image(query)
		fmt.Println("go image end")
	}()
	go func() {
		fmt.Println("go video")
		c <- Video(query)
		fmt.Println("go video end")
	}()

	var results []Result
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timer:
			return results, errors.New("timed out")
		}
	}

	return results, nil
}
