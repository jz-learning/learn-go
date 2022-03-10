package main

import "fmt"

func main() {
	i := 5

	var j string
	j = "hello"
	j += " world"

	// Can't convert automatically
	// j += i

	j += string(i)

	fmt.Println(j)
}
