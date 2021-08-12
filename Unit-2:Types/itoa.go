package main

import (
	"fmt"
	"strconv"
)

func main() {
	exStrconv()
	exSprintF()
	changeWithErr()
}

func exStrconv() {
	var (
		countdown = 10
		str       = "Launch in T minus " + strconv.Itoa(countdown) + " seconds."
	)
	fmt.Println(str)
}

func exSprintF() {
	var (
		countdown = 10
		str       = fmt.Sprintf("Launch in T minus %v seconds", countdown)
	)
	fmt.Println(str)
}

func changeWithErr() { //Change from string to int with err catch
	countdown, err := strconv.Atoi("10")
	if err != nil {
		fmt.Println("An error has occurred during type conversion")
	}
	fmt.Printf("Launch in T minus %d seconds", countdown)
}
