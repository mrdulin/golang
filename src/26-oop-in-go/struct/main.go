package main

import (
	"fmt"

	entities "github.com/mrdulin/golang/src/26-oop-in-go/struct/entities"
)

func main() {
	var tom entities.Person
	tom.Name, tom.Age = "Tom", 18
	bob := entities.Person{Age: 25, Name: "Bob"}
	paul := entities.Person{Name: "Paul", Age: 43}

	tbOlder, tbDiff := entities.Older(tom, bob)
	tpOlder, tpDiff := entities.Older(tom, paul)
	bpOlder, bpDiff := entities.Older(bob, paul)

	fmt.Printf("%s and %s, %s is older by %d years\n", tom.Name, bob.Name, tbOlder, tbDiff)
	fmt.Printf("%s and %s, %s is older by %d years\n", tom.Name, paul.Name, tpOlder, tpDiff)
	fmt.Printf("%s and %s, %s is older by %d years\n", bob.Name, paul.Name, bpOlder, bpDiff)
}
