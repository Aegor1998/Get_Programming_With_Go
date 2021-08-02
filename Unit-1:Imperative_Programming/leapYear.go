package main

import "fmt"

func main() {
	var year int

	fmt.Println("This program computes leap years.")
	fmt.Print("What year would you like to check: ")
	fmt.Scanln(&year)

	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)

	if leap { //checks if true
		fmt.Printf("Year %d is a leap year\n", year)
	} else {
		fmt.Printf("Year %d is not a leap year\n", year)
	}
	fmt.Println("Program is exiting")
}
