package main

import "fmt"

func main() {

	var j string
	j = "hello"
	j += " world"
	i := 5

	// Can't convert automatically
	// j += i

	j += fmt.Sprint(i)

	fmt.Println(j)

	// Can't convert int to double automatically
	velocity := 100 // int
	force := 2.5    // float64

	power := float64(velocity) * force

	fmt.Println(power)

	// Same thing but without conversion
	var velocity2 float64 = 5
	var force2 float64 = 1.5
	var power2 float64

	power2 = velocity2 * force2

	fmt.Println(power2)
}
