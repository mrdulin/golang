package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	files, _ := filepath.Glob("[")
	fmt.Println("matched files", files)
}
