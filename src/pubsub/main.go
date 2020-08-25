package main

import (
  "fmt"
  "github.com/docker/docker/pkg/pubsub"
  "strings"
  "time"
)

func main() {
  p := pubsub.NewPublisher(100*time.Microsecond, 10)
  golang := p.SubscribeTopic(func(v interface{}) bool {
    if key, ok := v.(string); ok {
      if strings.HasPrefix(key, "golang:") {
        return true
      }
    }
    return false
  })
  
  go p.Publish("hi")
  go p.Publish("golang: https://golang.org")
  fmt.Println("golang topic: ", <-golang)
}
  