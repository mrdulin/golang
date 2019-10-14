package main

import (
	"fmt"
	"reflect"
)

type googleAccount struct {
	ID           string
	RefreshToken string
}

func main() {
	googleAccount1 := &googleAccount{ID: "1", RefreshToken: "abc"}
	googleAccount2 := &googleAccount{ID: "2", RefreshToken: "xyz"}
	googleAccounts := []*googleAccount{googleAccount1, googleAccount2}
	fmt.Printf("googleAccounts=%v\n", googleAccounts)
	rval := checkPointerType(googleAccounts)
	fmt.Printf("rval=%v\n", rval)
	fmt.Printf("googleAccounts is pointer = %v\n", isPointerType(googleAccounts))

	googleAccounts2 := []googleAccount{
		{ID: "3", RefreshToken: "aaa"},
		{ID: "4", RefreshToken: "bbb"},
	}

	fmt.Printf("googleAccounts2=%v\n", googleAccounts2)
	rval2 := make([]bool, len(googleAccounts2))
	for i, v := range googleAccounts2 {
		rval2[i] = isPointerType(v)
	}
	fmt.Printf("rval2=%v\n", rval2)
	fmt.Printf("googleAccounts2 is pointer = %v\n", isPointerType(googleAccounts2))

	googleAccounts3 := &googleAccounts2
	fmt.Printf("googleAccounts3=%v\n", googleAccounts3)
	rval3 := make([]bool, len(googleAccounts2))
	for i, v := range *googleAccounts3 {
		rval2[i] = isPointerType(v)
	}
	fmt.Printf("rval3=%v\n", rval3)
	fmt.Printf("googleAccounts3 is pointer = %v\n", isPointerType(googleAccounts3))
}

func isPointerType(v interface{}) bool {
	return reflect.ValueOf(v).Kind() == reflect.Ptr
}

func checkPointerType(s []*googleAccount) []bool {

	rval := make([]bool, len(s))
	for i, v := range s {
		rval[i] = isPointerType(v)
	}
	return rval
}
