package main

import (
	"fmt"
)

func main() {
	var (
		message    = "uv vagreangvbany fcnpr fgngvba"
		newMessage string
	)

	for i := 0; i < len(message); i++ {
		temp := message[i]
		if temp >= 'a' && temp <= 'z' {
			temp += 13
			if temp > 'z' {
				temp -= 26
			}
		}
		newMessage += string(temp)

	}

	fmt.Printf("The cyphered message was %s\n", message)
	fmt.Printf("The message reads as %s\n", newMessage)
}
