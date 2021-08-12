package main

import "fmt"

func main() {
	var (
		age       = 23
		marsAge   = float64(age)
		marsDays  = 687.0
		earthDays = 365.2425
	)
	marsAge = marsAge * earthDays / marsDays
	fmt.Printf("I am %v Years old on Mars", int(marsAge))
}
