package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

const Url string = "http://localhost:9090"

func request(echo string) (string, error) {
	var rval string
	req, err := http.NewRequest("GET", Url, nil)
	if err != nil {
		return rval, err
	}

	client := &http.Client{}
	q := req.URL.Query()
	q.Add("echo", echo)
	req.URL.RawQuery = q.Encode()

	log.Println(req.URL.String())

	resp, err := client.Do(req)
	if err != nil {
		return rval, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return rval, err
	}
	rval = string(body)
	return rval, nil
}

func goroutineRequestWithChannel(echo string, ch chan string) {
	rval, _ := request(echo)
	ch <- rval
}

func goroutineRequestWithReturnVal(echo string) (interface{}, error) {
	return request(echo)
}

func sequence() {
	rval1, _ := request("aaa")
	log.Println("rval1: ", rval1)
	rval2, _ := request("bbb")
	log.Println("rval2: ", rval2)
	rval3, _ := request("ccc")
	log.Println("rval3: ", rval3)
}

func sequenceWithForLoop() {
	echos := []string{"aaa", "bbb", "ccc"}
	for _, echo := range echos {
		rval, _ := request(echo)
		log.Println("rval: ", rval)
	}
}

func concurrencyWithChannel() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	go goroutineRequestWithChannel("aaa", ch1)
	go goroutineRequestWithChannel("bbb", ch2)
	go goroutineRequestWithChannel("ccc", ch3)

	rval1, rval2, rval3 := <-ch1, <-ch2, <-ch3
	log.Printf("rval1: %s, rval2: %s, rval3: %s", rval1, rval2, rval3)
}

func concurrencyWithReturnVal() {

}

func main() {
	//sequence()
	//sequenceWithForLoop()
	//concurrencyWithChannel()
}
