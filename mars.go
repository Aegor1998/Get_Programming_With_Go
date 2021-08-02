package main

import (
	"fmt"
)

func main() {
	//unformatted()
	//formatted()
	//earth()
	testTable()
	//testTime()
	//variableSpeed()
}

func unformatted() {
	fmt.Print("My weight the surface of on Mars is ")
	//Mars is 0.3783 times the gravity of earth
	fmt.Print(235 * 0.3783)
	fmt.Print(" lbs, and I would be ")
	// 1 year on Earth is 365 days and 1 year on Mars is 687 days
	fmt.Print(23 * 365 / 687)
	fmt.Print(" years old.\n")
}

func formatted() { //Compresses the number of lines
	fmt.Printf("My weight on the surface of Mars is %v lbs,", 235*0.3783)
	fmt.Printf(" and I would be %v years old.\n", 23*365/687)
}

func earth() {
	fmt.Printf("My weight on the surface of %v is %v lbs.\n", "Earth", 235)
}

func testTable() {
	//When you place a number between % and v you will pad the number of spaces
	//to the left (+) or left (-)
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
}

func testTime() {
	const lightSpeed = 299792 //km/s
	var distance = 56000000   // km

	fmt.Println(distance/lightSpeed, " seconds")

	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")
}

func variableSpeed() {
	var (
		distance = 96300000 //km between Earth and Mars
		speed    = 100800   // km/h
		hours    = distance / speed
		days     = hours / 24
		date     = "January 2025" //launch Date
	)

	fmt.Printf("If SpaceX launches a flight to mars in %v, going %v", date, speed)
	fmt.Printf(", it will take them %v days to get there.\n", days)
}
