package main

import (
	"fmt"
	"time"
)

func doSomething() {
	fmt.Println("do something")
}

// 结果可能是先打印do something，后打印main function, 但想象一下，这里如果把fmt.Println("do something")改为某个执行时间和休眠时间相同的操作
// 那么执行顺序就不确定了，所以goroutine之间要保证执行顺序，必须结合通道和sync来进行同步和排序
func main() {
	go doSomething()

	time.Sleep(1 * time.Second)
	fmt.Println("main function")
}
