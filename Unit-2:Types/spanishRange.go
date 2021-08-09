package main

import "fmt"

func main() {
	question := "¿Cómo estás?"

	for index, cypher := range question {
		fmt.Printf("%v %c %d\n", index, cypher, cypher)
	}
}
