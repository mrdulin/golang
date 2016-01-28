package main

import (
	"fmt"
)

func sendData(ch chan<- int) {
	ch <- 666
}

func main() {
	sendOnlyCh := make(chan<- int)
	go sendData(sendOnlyCh)
	// 由于sendOnlyCh是单向的，只能向sendOnlyCh发送数据，无法从sendOnlyCh读取数据
	// 因此报错: invalid operation: <-sendOnlyCh (receive from send-only type chan<- int)
	// fmt.Println(<-sendOnlyCh)

	// 将双向channel转换为单向channel，反之亦然
	ch := make(chan int)
	go sendData(ch)
	fmt.Println(<-ch)
}
