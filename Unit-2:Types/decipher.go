package main

import "fmt"

func main() {
	const (
		KEYWORD    string = "GOLANG"
		CYPHERTEXT string = "CSOITEUIWUIZNSROCNKFD"
	)
	var (
		decodedText string
		cypherKey   = 0
	)

	fmt.Println("Program Start")

	for i := 0; i < len(CYPHERTEXT); i++ {
		fmt.Println("Loop # ", i)
		if cypherKey == len(KEYWORD) {
			cypherKey = 0
		}
		decodedText += letterDecoder(rune(CYPHERTEXT[i]), rune(KEYWORD[cypherKey]))
		cypherKey++
	}
	fmt.Println(decodedText)
}

//Using a switch function this takes a character from the keyword and turns it into the corresponding value
func letterDecoder(cipherLetter rune, keyLetter rune) (letter string) {
	mapkey := map[rune]int{
		'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4, 'F': 5,
		'G': 6, 'H': 7, 'I': 8, 'J': 9, 'K': 10, 'L': 11,
		'M': 12, 'N': 13, 'O': 14, 'P': 15, 'Q': 16, 'R': 17,
		'S': 18, 'T': 19, 'U': 20, 'V': 21, 'W': 22, 'X': 23,
		'Y': 24, 'Z': 25,
	}
	letterValue := mapkey[cipherLetter] - mapkey[keyLetter]
	if letterValue < 0 {
		letterValue += 26
	}
	for keys, value := range mapkey {
		if value == letterValue {
			letter = string(keys)
		}
	}
	return letter
}
