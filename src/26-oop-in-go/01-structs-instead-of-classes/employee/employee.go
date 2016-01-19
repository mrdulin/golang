package employee

import "fmt"

type Employee struct {
	FirstName, LastName      string
	TotalLeaves, LeavesTaken int
}

func (e Employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
