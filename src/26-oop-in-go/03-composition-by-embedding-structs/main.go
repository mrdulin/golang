package main

import "github.com/mrdulin/golang/src/26-oop-in-go/03-composition-by-embedding-structs/models"

func main() {
	author1 := models.Author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}
	post1 := models.Post{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	post1.Details()
}
