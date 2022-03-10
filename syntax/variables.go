package main

import "fmt"

// declaration, need type
var c, python, java bool

// declaration + initialization, no need type
var i, j = 5, 10

func someMath(x int) int {
	// short hand declaration, no need type
	multiple := 10
	return x * multiple
}

func main() {
	var i int
	var tall, short, name = true, false, "bob"

	fmt.Println(name, "is tall:", tall, "and is short:", short)
	fmt.Println(i, j, c, python, java)
	fmt.Println(someMath(5))

}
