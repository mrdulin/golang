package main

import (
	"sync"
)

// func main() {
//   go println("hello world")
// }

// goroutine和main函数的返回是并发的，谁都有可能先发生
// ☁  golang [master] ⚡  go run /Users/ldu020/workspace/github.com/mrdulin/golang/src/goroutine/00-hello-world/main.go
// hello world%
// ☁  golang [master] ⚡
// ☁  golang [master] ⚡  go run /Users/ldu020/workspace/github.com/mrdulin/golang/src/goroutine/00-hello-world/main.go
// hello world%
// ☁  golang [master] ⚡  go run /Users/ldu020/workspace/github.com/mrdulin/golang/src/goroutine/00-hello-world/main.go
// hello world%
// ☁  golang [master] ⚡  go run /Users/ldu020/workspace/github.com/mrdulin/golang/src/goroutine/00-hello-world/main.go
// ☁  golang [master] ⚡  go run /Users/ldu020/workspace/github.com/mrdulin/golang/src/goroutine/00-hello-world/main.go
// hello world%

// 解决方法1:
// 通过同步原语来保证执行顺序, <-done执行时，必然是done <- 1已经执行, 同一个goroutine满足顺序一致性规则，即当 done <- 1执行时， println("hello, world")语句必然已经执行完成
// func main() {
// 	done := make(chan int)
// 	go func() {
// 		println("hello, world")
// 		done <- 1
// 	}()
// 	<-done
// }

// 解决方法2:
func main() {
	var mu sync.Mutex
	mu.Lock()
	go func() {
		println("hello, world")
		mu.Unlock()
	}()
	mu.Lock()
}
