package main

import (
	"fmt"
	"math"
)

type gpsGPS struct {
	worldGPS
	currentLocation     locationGPS
	destinationLocation locationGPS
}

type worldGPS struct {
	name   string
	radius float64
}

type locationGPS struct {
	name string
	long float64
	lat  float64
}

type roverGPS struct {
	name string
	gpsGPS
}

func radConGPS(deg float64) float64 {
	return deg * math.Pi / 180
} //Coverts Degrees to Radians

func (w worldGPS) distance(p1, p2 locationGPS) float64 {
	var (
		s1, c1 = math.Sincos(radConGPS(p1.lat))
		s2, c2 = math.Sincos(radConGPS(p2.lat))
		clong  = math.Cos(radConGPS(p1.long - p2.long))
	)

	return w.radius * math.Acos(s1*s2+c1*c2*clong)
} //Returns distance between 2 points on same planet

func main() {
	var (
		mars          = worldGPS{name: "Mars", radius: 3396.2} //Source https://nssdc.gsfc.nasa.gov/planetary/factsheet/marsfact.html
		startLocation = locationGPS{name: "Bradbury_Landing", long: -45.895, lat: 137.4417}
		endLocation   = locationGPS{name: "Elysium_Planitia", long: 4.5, lat: 135.9}
		currentRoute  = gpsGPS{worldGPS: mars, currentLocation: startLocation, destinationLocation: endLocation}
		Curiosity     = roverGPS{name: "Curiosity", gpsGPS: currentRoute}
	)
	findDistance(Curiosity)
}

func findDistance(Curiosity roverGPS) {
	fmt.Printf("\nThe %s Rover is currently located at %s.\n", Curiosity.name, Curiosity.gpsGPS.currentLocation.name)
	fmt.Printf("The %s Rover is headed to %s.\n", Curiosity.name, Curiosity.gpsGPS.destinationLocation.name)
	fmt.Printf("The distance between the two locations is %v KM.\n", Curiosity.worldGPS.distance(Curiosity.currentLocation, Curiosity.destinationLocation))
}
