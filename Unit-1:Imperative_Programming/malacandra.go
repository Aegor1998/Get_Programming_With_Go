package main

import (
	"fmt"
)

func main() {
	fmt.Println("Program Start")

	const (
		days     = 28
		distance = 56000000 //km
		hours    = days * 24
		planet   = "malacandra"
	)
	var speed = distance / hours // km/hr

	fmt.Printf("If the time alloted is %v and the distance is %d", days, distance)
	fmt.Printf(" the ship will have to go %v km/hr to reach %v", speed, planet)
	fmt.Printf(" in time.\n")
}
