package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	name := "inanç           "
	fmt.Println(len(name))
	name = strings.TrimRight(name, " ")
	fmt.Println(name)
	fmt.Println(utf8.RuneCountInString(name))
}
