package main

import "fmt"

// declaration, need type
var c, python, java bool

// declaration + initialization, no need type
var i, j = 5, 10

// naming convention
var time int
var Time int // will be exported
var _ int    // works as well

// int initializes as 0

func someMath(x int) int {
	// short hand declaration, no need var or type
	multiple := 10
	return x * multiple
}

func main() {
	var i int
	tall, short, name := true, false, "bob"

	// group variables together for readability
	var (
		age       int
		address   string
		birthDate string
	)

	fmt.Println(name, "is tall:", tall, "and is short:", short)
	fmt.Println(i, j, c, python, java)
	fmt.Println(someMath(5))

}
