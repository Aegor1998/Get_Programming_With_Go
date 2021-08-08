package main

import "fmt"

func main() {
	const (
		distance      = 236000000000000000                    // km
		lightSpeed    = 299792                                // km/s
		secondsPerDay = 86400                                 // s
		days          = distance / lightSpeed / secondsPerDay // km/day
	)
	fmt.Printf("The galaxy Canis Major Dwarf is %v days from out sun", days)
}
