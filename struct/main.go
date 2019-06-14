package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func main() {
	var tom person
	tom.name, tom.age = "Tom", 18
	bob := person{age: 25, name: "Bob"}
	paul := person{"Paul", 43}

	tbOlder, tbDiff := older(tom, bob)
	tpOlder, tpDiff := older(tom, paul)
	bpOlder, bpDiff := older(bob, paul)

	fmt.Println("%s and %s, %s is older by %d years\n", tom.name, bob.name, tbOlder, tbDiff)	
	fmt.Println("%s and %s, %s is older by %d years\n", tom.name, paul.name, tpOlder, tpDiff)
	fmt.Println("%s and %s, %s is older by %d years\n", bob.name, paul.name, bpOlder, bpDiff)
}
