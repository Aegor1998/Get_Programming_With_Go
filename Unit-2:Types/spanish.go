package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	question := "¿Cómo estás?"

	fmt.Printf("%s is %d bytes long\n", question, len(question))
	fmt.Printf("%s has %v runes in it\n", question, utf8.RuneCountInString(question))
	c, size := utf8.DecodeRuneInString(question)
	fmt.Printf("The first rune %c has %v bytes\n", c, size)
}
