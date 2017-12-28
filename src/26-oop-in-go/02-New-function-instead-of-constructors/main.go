package main

import "26-oop-in-go/02-New-function-instead-of-constructors/employee"

func main() {
	e := employee.New("Sam", "Adolf", 30, 20)
	e.LeavesRemaining()
}
