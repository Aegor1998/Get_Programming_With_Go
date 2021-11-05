package main

import (
	"fmt"
	"math"
)

type worldDis struct {
	name   string
	radius float64
}

type locationDis struct {
	name string
	long float64
	lat  float64
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
} //Coverts Degrees to Radians

func (w worldDis) distance(p1, p2 locationDis) float64 {
	var (
		s1, c1 = math.Sincos(rad(p1.lat))
		s2, c2 = math.Sincos(rad(p2.lat))
		clong  = math.Cos(rad(p1.long - p2.long))
	)

	return w.radius * math.Acos(s1*s2+c1*c2*clong)
} //Returns distance between 2 points on same planet

func main() {
	var (
		earth = worldDis{name: "Earth", radius: 6378.0}
	)
	disLondonFrance(earth)
	disToDC(earth)

}

func disLondonFrance(earth worldDis) {
	var (
		london = locationDis{name: "London", long: 51.5, lat: 0.13}
		paris  = locationDis{name: "Paris", long: 48.85, lat: 2.35}
		dist   = earth.distance(london, paris)
	)
	fmt.Printf("\n The distance between %s and %s is %v km.\n", london.name, paris.name, dist)
}

func disToDC(earth worldDis) {
	var (
		indy = locationDis{name: "Indy, IN", lat: 39.77758949810973, long: -86.15747454158341}
		dc   = locationDis{name: "Washington DC", lat: 38.91370754331705, long: -77.03630759024611}
		dist = earth.distance(indy, dc)
	)
	fmt.Printf("\n The distance between %s and %s is %v km.\n", indy.name, dc.name, dist)
}
