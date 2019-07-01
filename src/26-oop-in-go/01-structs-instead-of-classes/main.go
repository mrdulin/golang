package main

import "26-oop-in-go/01-structs-instead-of-classes/employee"

func main() {
  e := employee.Employee{
    FirstName: "Sam",
    LastName: "Adolf",
    TotalLeaves: 30,
    LeavesTaken: 20,
  }
  e.LeavesRemaining()

  e2 := employee.Employee{}
  e2.LeavesRemaining()
}