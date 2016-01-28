package main

import (
	"fmt"
)

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go producer(ch)
	// 无限循环
	// for {
	// 	v, ok := <-ch
	// 	// 检查channel是否被close, 如果是，ok为false，跳出循环
	// 	if ok == false {
	// 		break
	// 	}
	// 	fmt.Println("Received: ", v, ok)
	// }

	fmt.Println("for range style")
	// 也可以使用for range语法, 当channel被close后，会自动跳出循环
	for v := range ch {
		fmt.Println("Received: ", v)
	}
}
