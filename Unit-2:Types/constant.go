package main

import "fmt"

func main() {
	const (
		distance      = 24000000000000000000
		lightSpeed    = 299792
		secondsPerDay = 86400
		days          = distance / lightSpeed / secondsPerDay
		years         = days / 365
	)
	fmt.Printf("The Andromeda Galaxy is %v light years away", years)
}
