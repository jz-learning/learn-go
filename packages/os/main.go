package main

import (
	"fmt"
	"os"
)

func main() {
	// go run main.go hello james
	// Args[0] - path to the program
	// Args[1] - "hello"
	// Args[2] - "james"

	args := os.Args
	fmt.Printf("%#v\n", args)

	greeting := args[1]
	name := args[2]

	fmt.Println(greeting)
	fmt.Println(name)
}
