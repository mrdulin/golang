package models

import "fmt"

type Author struct {
	FirstName, LastName, Bio string
}

func (a Author) fullName() string {
	return fmt.Sprintf("%s %s", a.FirstName, a.LastName)
}
