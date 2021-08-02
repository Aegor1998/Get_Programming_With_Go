package main

import (
	"fmt"
)

func main() {
	var cont bool

	fmt.Print("Do you with to start the game(y/n): ")
	cont = userInput()

	if cont == true {
		fmt.Println("You are on a very large and dangerous mountain. You'll want to seek shelter to survive")
		fmt.Print("You spot a cave entrance do you enter(y/n): ")
		cont = userInput()
		if cont == true {
			fmt.Println("You arrive at the cave enterance. A storm is brewing outside")
			fmt.Print("Do you do deaper into the cave(y/n): ")
			cont = userInput()
			if cont == true {
				fmt.Println("You enter the cave and are eaten by a Grue!")
			} else {
				fmt.Println("The storm carries you off the mountain to safety.")
			}
		} else {
			fmt.Println("The land gives out under you and you fall to your death.")
		}
	}

	fmt.Println("Program is exiting.")
}

func userInput() bool {
	var input string
	fmt.Scan(&input)
	fmt.Println(input)
	if input == "y" {
		return true
	}
	return false
}
