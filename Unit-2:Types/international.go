package main

import "fmt"

func main() {
	const text string = "Hola Estación Espacial Internacional"
	var cypherText string

	for i := 0; i < len(text); i++ {
		cypherText += rot13Encrypt(rune(text[i]))
	}
	fmt.Printf("Text:   %s\n", text)
	fmt.Printf("Cypher: %s\n", cypherText)
}

//this is essentially rot13.go but using switch statements
func rot13Encrypt(character rune) string {
	switch {
	case character == 'ó': //This is evaluated at the correct point, but it seems that 'ó' == 'Ã³'
		return string(character)
	case character == ' ':
		return string(character)
	case character >= 'a' && character <= 'z':
		character += 13
		if character > 'z' {
			character -= 26
		}
	}
	return string(character)
}
