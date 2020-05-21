package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(i int, ctx context.Context, wg *sync.WaitGroup) error {
  defer wg.Done()
  for {
    select {
    default:
      fmt.Printf("worker: %d hello\n", i)
    case <- ctx.Done():
      fmt.Printf("worker: %d exited\n", i)
      return ctx.Err()
    }
  }
}

func main() {
  var wg sync.WaitGroup
  ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
  for i := 0; i < 10; i++ {
    wg.Add(1)
    go worker(i, ctx, &wg)
  }
  time.Sleep(time.Second)
  cancel()
  wg.Wait()
}