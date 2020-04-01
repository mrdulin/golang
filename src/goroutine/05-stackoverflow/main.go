package main

import (
  "fmt"
  "runtime"
)

func consumer(ch chan int) {
	for {
    data := <-ch
    
    if data == 0 {
      break
    }

    fmt.Println(data)
  }
  fmt.Println("goroutine exit")
}

func main() {
  ch := make(chan int)
  for {
    var dummy string
    fmt.Scan(&dummy)
    if dummy == "quit" {
      for i := 0; i < runtime.NumGoroutine() - 1; i ++ {
        ch <- 0
      }
      continue
    }
    go consumer(ch)
    fmt.Println("goroutines: ", runtime.NumGoroutine())
  }
}