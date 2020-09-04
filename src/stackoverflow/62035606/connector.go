package connector

import (
	"fmt"
)

type Pinger interface {
	Ping() error
}

type Connector struct {
	DB Pinger
}

func (c *Connector) Pool() (interface{}, error) {
	err := c.DB.Ping()
	if err != nil {
		fmt.Println("error handle logic")
		return nil, err
	}
	fmt.Println("success logic")
	return 1, nil
}
