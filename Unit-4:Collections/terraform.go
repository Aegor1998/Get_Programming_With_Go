package main

import "fmt"

type planet string

func (s planet) terraform() planet {
	if s == "Earth" {
		return s + "-Prime"
	}
	return "New-" + s
} //appends "New-" to a planet name

func main() {
	solarSystem := [8]planet{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	terrestrial := solarSystem[0:4]
	gasGiants := solarSystem[4:6]
	iceGiants := solarSystem[6:8]

	fmt.Println(terrestrial, gasGiants, iceGiants)
	for i := 0; i < len(solarSystem); i++ {
		solarSystem[i] = solarSystem[i].terraform()
	}
	terrestrial = solarSystem[0:4]
	gasGiants = solarSystem[4:6]
	iceGiants = solarSystem[6:8]
	fmt.Println(terrestrial, gasGiants, iceGiants)
}
