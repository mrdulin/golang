package models 

import "fmt"

type Post struct {
  Title, Content string
  Author Author
}

func (p Post) Details() {  
  fmt.Println("Title: ", p.Title)
  fmt.Println("Content: ", p.Content)
  fmt.Println("Author: ", p.Author.fullName())
  fmt.Println("Bio: ", p.Author.Bio)
}