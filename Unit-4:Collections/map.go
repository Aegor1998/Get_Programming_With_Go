package main

import "fmt"

func main() {
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}
	temp := temperature["Earth"]
	fmt.Printf("On average Earth is %v C.\n", temp)

	temperature["Earth"] = 16
	temperature["Venus"] = 464

	fmt.Println(temperature)

	if moon, ok := temperature["Moon"]; ok { //don't have to use ok. Could make a variable on any name of type bool
		fmt.Printf("On average Moon is %v C.\n", moon)
	} else {
		fmt.Println("Where is the Moon?")
	}
}
