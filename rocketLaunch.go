package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var (
		count int = 100
	)

	for count > 0 {
		fmt.Printf("Count: %v\n", count)
		if rand.Intn(count) == 100%count {
			break
		}
		count--
	}
	if count == 0 {
		fmt.Println("Launch was successful.")
	} else {
		fmt.Println("Launch Failed.")
	}
}
