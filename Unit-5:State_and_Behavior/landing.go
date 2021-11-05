//Start Monday, October 18th, 2021 13:29 (Blake Kucera)
//  End Monday, October 18th, 2021 13:45 (Blake Kucera)
//Experiment: landing.go
/*
Write a program that displays the JSON encoding of the three roverGPS landing sites in listing 21.8.
	The JSON should include the name of each landing site and use struct tags as shown in listing 21.10.
To make the output friendlier, make use of the MarshalIndent function from the json package.
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type locationLanding struct {
	Rover       string  `json:"roverGPS"`
	LandingSite string  `json:"landing-site"`
	Lat         float64 `json:"lat"`
	Long        float64 `json:"long"`
}

func main() {
	locations := []locationLanding{
		{Rover: "Curiosity", LandingSite: "Bradbury_Landing", Lat: -4.5895, Long: 137.4417},
		{Rover: "Spirit", LandingSite: "Columbia_Memorial_Station", Lat: -14.5684, Long: 175.472636},
		{Rover: "Opportunity", LandingSite: "Challenger_Memorial_Station", Lat: -1.9462, Long: 354.4734},
		{Rover: "InSight", LandingSite: "Elysium_Planitia", Lat: 4.5, Long: 135.9},
	}
	bytes, err := json.MarshalIndent(locations, "", "")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(bytes))

	for i := range locations {
		fmt.Printf("Rover %v landed at %v, Location: LAT:%v, LONG:%v\n", locations[i].Rover, locations[i].LandingSite, locations[i].Lat, locations[i].Long)
	}
}
