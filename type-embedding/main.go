package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

type admin struct {
	user
	level string
}

func (u *user) notify() {
	fmt.Printf("sending user email to %s<%s>\n", u.name, u.email)
}

func (a *admin) notify() {
	fmt.Printf("sending admin user to %s<%s>\n", a.name, a.email)
}

func main() {
	ad := admin{
		user: user{
			name:  "mrdulin",
			email: "novaline.dulin@gmail.com",
		},
		level: "super",
	}

	ad.notify()
	ad.user.notify()

	sendNotification(&ad)
}

func sendNotification(n notifier) {
	n.notify()
}
