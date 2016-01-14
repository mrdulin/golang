package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type IStore interface {
	search(name string) (string, error)
}

type BookStore struct {
	apiURL     *string
	httpClient *http.Client
}

func NewBookStore(apiURL *string, httpClient *http.Client) IStore {
	return &BookStore{
		apiURL,
		httpClient,
	}
}

func (bookStore *BookStore) search(name string) (string, error) {
	url := *bookStore.apiURL + "/search/" + name
	resp, err := bookStore.httpClient.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func main() {
	var apiURL = "https://api.itbook.store/1.0"
	httpClient := &http.Client{}
	bookStore := NewBookStore(&apiURL, httpClient)

	var wg sync.WaitGroup
	var bookNames = [3]string{"golang", "typescript", "nodejs"}
	var rvals = make([]string, len(bookNames))
	for _, bookName := range bookNames {
		wg.Add(1)
		go func(bookName string) {
			defer wg.Done()
			rval, _ := bookStore.search(bookName)
			rvals = append(rvals, rval+"\n")
		}(bookName)
	}
	wg.Wait()

	fmt.Println("rvals: ", rvals)
}
