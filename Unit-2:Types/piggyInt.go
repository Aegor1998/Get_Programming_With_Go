package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	nickle  int = 5
	dime    int = 10
	quarter int = 25
)

var (
	piggyBank      int
	nickleCounter  int
	dimeCounter    int
	quarterCounter int
)

func main() {
	for limit := false; limit == false; {
		chooseCoin()
		limit = limitReached()
		fmt.Printf("Piggy Bank Balance: $%.2f\n", float64(piggyBank)/100)
	}
	printContents()
}

func chooseCoin() {
	rand.Seed(time.Now().UnixNano() * time.Now().UnixNano())
	choice := rand.Intn(4) - 1
	switch choice {
	case 0:
		piggyBank += nickle
		nickleCounter++
	case 1:
		piggyBank += dime
		dimeCounter++
	case 2:
		piggyBank += quarter
		quarterCounter++
	}
}

func limitReached() bool {
	if piggyBank >= 2000 { //This is 2000 cents or $20.00
		return true
	} else {
		return false
	}
}

func printContents() {
	fmt.Println("The Piggy Bank has reach its target balance")
	fmt.Printf("Piggy Bank Balance: $%.2f\n", float64(piggyBank)/100)
	fmt.Printf("Nickles : %d\n", nickleCounter)
	fmt.Printf("Dimes   : %d\n", dimeCounter)
	fmt.Printf("Quarters: %d\n", quarterCounter)
}
