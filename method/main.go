package main

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

func (user user) notify() {
	fmt.Printf("sending user email to %s<%s>\n", user.name, user.email)
}

func (user *user) changeEmail(email string) {
	user.email = email
}

func main() {
	bill := user{name: "Bill", email: "bill@gmail.com"}
	bill.notify()

	lisa := &user{"Lisa", "lisa@gmail.com"}
	lisa.notify()

	bill.changeEmail("bill@qq.com")
	bill.notify()

	lisa.changeEmail("lisa@qqq.com")
	lisa.notify()
}
