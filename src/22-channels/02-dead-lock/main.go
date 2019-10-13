package main

import "fmt"

// cannot recover from a deadlock.
// https://stackoverflow.com/questions/14075382/how-do-i-catch-the-exception-of-a-channel-deadlocking
// https://stackoverflow.com/questions/16783040/how-to-recover-the-deadlock-error-in-golang?noredirect=1&lq=1
// deadlock is not a panic error, it's a fatal error, see https://golang.org/pkg/log/#Fatal
func main() {
	defer func() {
		r := recover()
		fmt.Printf("rrrr=%v\n", r)
		if r != nil {
			fmt.Printf("r=%v\n", r)
		}
	}()

	ch := make(chan int)
	ch <- 5
	fmt.Println("code will never arrive here")
}
