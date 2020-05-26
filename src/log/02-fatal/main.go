package main

import (
  "log"
  "fmt"
)

func ping() bool {
  log.Fatal("acquire conn fail")
  return false
}

func main() {
  pong := ping()
  fmt.Print(pong)
}