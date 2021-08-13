package main

import (
	"fmt"
	"strings"
)

const KEYWORD string = "GOLANG"

func main() {
	var (
		choice int
	)
	fmt.Printf("Do you wish to encode(0) or decode(1): ")
	fmt.Scan(&choice)

	switch choice {
	case 0: //encode
		fmt.Println("You Chose ")
		encodeMessage()
	case 1: //decode
		decodeMessage()
	}
}

func getMessage() (message string) {
	fmt.Println("Do not put in symbols or spaces")
	fmt.Scan(&message)
	return strings.ToUpper(message)
}

func encodeMessage() {
	var (
		message         string
		encodedMessage  string
		keywordPosition = 0
	)

	fmt.Println("You have entered the Message Encoder.")
	//Function Logic
	message = getMessage()
	//Start Encoding Message
	for i := 0; i < len(message); i++ {
		if keywordPosition == len(KEYWORD) {
			keywordPosition = 0
		} //Resets keywordPosition
		encodedMessage += messageEncoder(rune(message[i]), rune(KEYWORD[keywordPosition]))
		keywordPosition++
	}
	fmt.Printf("Message:         %s\n", message)
	fmt.Printf("Keyword:         %s\n", KEYWORD)
	fmt.Printf("Encoded Message: %s\n", encodedMessage)
}

func messageEncoder(cipherLetter rune, keyLetter rune) (letter string) {
	mapkey := map[rune]int{
		'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4, 'F': 5,
		'G': 6, 'H': 7, 'I': 8, 'J': 9, 'K': 10, 'L': 11,
		'M': 12, 'N': 13, 'O': 14, 'P': 15, 'Q': 16, 'R': 17,
		'S': 18, 'T': 19, 'U': 20, 'V': 21, 'W': 22, 'X': 23,
		'Y': 24, 'Z': 25,
	}
	letterValue := mapkey[cipherLetter] + mapkey[keyLetter]
	if letterValue > 25 {
		letterValue -= 25
	}
	for keys, value := range mapkey {
		if value == letterValue {
			letter = string(keys)
		}
	}
	return letter
}

func decodeMessage() {
	var (
		message         string
		decodedMessage  string
		keywordPosition = 0
	)
	fmt.Println("You have entered the Message Decoder.")
	message = getMessage()

	for i := 0; i < len(message); i++ {
		if keywordPosition == len(KEYWORD) {
			keywordPosition = 0
		} //Resets keywordPosition
		decodedMessage += messageDecoder(rune(message[i]), rune(KEYWORD[keywordPosition]))
		keywordPosition++
	}
	fmt.Printf("Message:         %s\n", message)
	fmt.Printf("Keyword:         %s\n", KEYWORD)
	fmt.Printf("Decoded Message: %s\n", decodedMessage)
}

func messageDecoder(cipherLetter rune, keyLetter rune) (letter string) {
	mapkey := map[rune]int{
		'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4, 'F': 5,
		'G': 6, 'H': 7, 'I': 8, 'J': 9, 'K': 10, 'L': 11,
		'M': 12, 'N': 13, 'O': 14, 'P': 15, 'Q': 16, 'R': 17,
		'S': 18, 'T': 19, 'U': 20, 'V': 21, 'W': 22, 'X': 23,
		'Y': 24, 'Z': 25,
	}
	fmt.Printf("cipherletter: %s, keyletter: %s\n", string(cipherLetter), string(keyLetter))
	letterValue := mapkey[cipherLetter] + mapkey[keyLetter]
	if letterValue > 25 {
		letterValue -= 25
	}
	for keys, value := range mapkey {
		if value == letterValue {
			letter = string(keys)
		}
	}
	return letter
}
