package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	nickleValue  = 0.05
	dimeValue    = 0.10
	quarterValue = 0.25
)

func main() {
	var (
		piggyBalance   float64
		nickleCounter  int
		dimeCounter    int
		quarterCounter int
	)
	for target := false; target == false; {
		coin := whatCoin()
		piggyBalance += coin
		switch coin {
		case nickleValue:
			nickleCounter += 1
		case dimeValue:
			dimeCounter += 1
		case quarterValue:
			quarterCounter += 1
		}
		displayDollars(piggyBalance)
		piggySize(piggyBalance, &target)
	}
	results(piggyBalance, nickleCounter, dimeCounter, quarterCounter)
}

//uses RAND to choose a coin
func whatCoin() (value float64) {
	rand.Seed(time.Now().UnixNano())
	coinChoice := rand.Intn(3)
	switch coinChoice {
	case 0:
		value = nickleValue
	case 1:
		value = dimeValue
	case 2:
		value = quarterValue
	}
	return value
}

//Checks to see if the Piggy Bank has reached $20.00 yet.
func piggySize(piggyBalance float64, target *bool) {
	if piggyBalance >= 20 {
		*target = true
	} else {
		*target = false
	}
}

//formats the display of piggyBalance
func displayDollars(piggyBalance float64) {
	fmt.Printf("Piggy Bank Balance = $%.2f\n", piggyBalance)
}

//Displays Results
func results(piggyBalance float64, nickleCounter int, dimeCounter int, quarterCounter int) {
	fmt.Println("Target Balance Has Been Reached")
	displayDollars(piggyBalance)
	fmt.Printf("Nickels: %d\n", nickleCounter)
	fmt.Printf("Dimes: %d\n", dimeCounter)
	fmt.Printf("Quarters: %d\n", quarterCounter)
}
