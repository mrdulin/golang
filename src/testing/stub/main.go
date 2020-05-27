package main

import (
  "fmt"
)

// https://stackoverflow.com/questions/62035606/how-to-stub-method-called-in-another-method-in-same-struct
// https://stackoverflow.com/a/47643967/6463558
type Connector struct {}

func (c *Connector) Pool() (interface{}, error) {
  err := c.ping()
  if err != nil {
    fmt.Println("error handle logic");
    return nil, err
  }
  fmt.Println("success logic")
  return 1, nil
}

func (c *Connector) ping() error {
  var err error
  // err = some side-effect RPC operation
  if err != nil {
    return err
  }
  return nil
}