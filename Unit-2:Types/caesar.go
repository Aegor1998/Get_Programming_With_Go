package main

import "fmt"

func main() {
	const cypher uint8 = 3
	var (
		message    = "L fdph, L vdz, L frqtxhuhg"
		newMessage string
	)

	for i := 0; i < len(message); i++ {
		if message[i] != ' ' && message[i] != ',' {
			newMessage += string(message[i] - cypher)
		} else {
			newMessage += string(message[i])
		}
	}

	fmt.Printf("The cyphered message was %s\n", message)
	fmt.Printf("The message reads as %s\n", newMessage)
}
