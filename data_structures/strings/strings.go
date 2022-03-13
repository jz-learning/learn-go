package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func rawString() {
	// Regular old strings
	a := "<body>\n\t<p>\"hello world\"</p>\n</body>"

	// Raw String literals
	b := `
<body>
	<p>"hello world"</p>
</body>`

	println(a)
	println(b)
}

func stringLength() {
	// length of string
	name := "résumé"

	// Except len doesn't return how many characters, but how many bytes
	// Each character can be 1 byte or 4 bytes
	length := len(name)
	println(length)

	// Runes can count in units of 4 bytes
	nRunes := utf8.RuneCountInString(name)
	println(nRunes)
}

// banger returns a string that is the uppercase version of the parameter
// followed by the same number of exclamation mark as the length of the param
func banger(s string) string {
	n := utf8.RuneCountInString(s)
	exclamtion := strings.Repeat("!", n)
	upper := strings.ToUpper(s)

	return exclamtion + upper + exclamtion
}

func main() {
	rawString()
	stringLength()
	fmt.Println(banger(os.Args[1]))
}
