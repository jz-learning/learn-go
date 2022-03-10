package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func add(x, y int) int {
	return x + y
}

func diffType(num int, name string) {
	fmt.Printf("there are %d %s(s)\n", num, name)
}

func namedReturn(num int) (double, triple int) {
	double = num * 2
	triple = num * 3

	// return double, triple
	return
}

func main() {
	fmt.Println(swap("hello", "world"))

	diffType(5, "james")

	fmt.Println(namedReturn(10))
}
