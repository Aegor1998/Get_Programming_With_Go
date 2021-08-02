package main

import "fmt"

func main() {
	var room = "lake"

	switch room {
	case "cave":
		fmt.Println("You are mostly certain that this is a cave")
	case "lake":
		fmt.Println("The frozen lake feels solid enough")
		fallthrough
	case "underwater":
		fmt.Println("The water is freezing cold.")
	}
}
