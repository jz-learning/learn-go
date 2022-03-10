package main

import (
	"fmt"
	"path"
)

func main() {
	// returns two values, the directory and file name
	dir, file := path.Split("example/text.txt")
	fmt.Println(dir)
	fmt.Println(file)
}
