package main

import (
	"fmt"
)

func main() {
	name := "Naveen"
	fmt.Printf("Original String: %s\n", string(name))
	fmt.Printf("Reversed String: ")
	for _, v := range []rune(name) {
		// defer函数调用栈，将以后进先出的顺序执行
		defer fmt.Printf("%c", v)
	}
}
