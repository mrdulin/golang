package employee

import "fmt"

// 如果不导出struct，一般也将struct中的字段不导出，都改成小写字母开头
type employee struct {
  firstName, lastName string
  totalLeaves, leavesTaken int
}

// New 一般使用New函数作为类似Java中constructor, 如果一个package中只有一个struct，用New函数，如果有多个struct，则用NewT命名，例如NewEmployee
func New(firstName string, lastName string, totalLeaves int, leavesTaken int) employee {
  return employee{firstName, lastName, totalLeaves, leavesTaken}
}

func (e employee) LeavesRemaining() {  
  fmt.Printf("%s %s has %d leaves remaining", e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
}