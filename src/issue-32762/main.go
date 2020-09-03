package main

import (
	"fmt"

	"github.com/mrdulin/golang/src/issue-32762/domain/models"
)

func main() {
	user := models.User{UserID: 1}

	fmt.Printf("user=%v\n", user)
}
