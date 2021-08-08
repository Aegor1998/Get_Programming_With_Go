package main

import "fmt"

func main() {
	const (
		lightSpeed    int = 299792 // km/s
		secondsPerDay int = 86400
	)
	var (
		distance int64 = 41.3e12 //distance to Alpha Centaur
	)
	fmt.Printf("Alpha Centaur is %d km away.\n", distance)

	daysToTravel := distance / int64(lightSpeed*secondsPerDay)

	fmt.Printf("That is %d days of traveling at light speed.\n", daysToTravel)

}
