/*
Goal: Create a program that outputs coordinates in a JSON having decimal degrees and degrees, minutes, seconds format
	Example:
		{
    "decimal": 135.9,
    "dms": "135°54'0.0\" E",
    "degrees": 135,
    "minutes": 54,
    "seconds": 0,
    "hemisphere": "E"
		}
*/
package main

import (
	"fmt"
	"strconv"
	"unicode"
)

type coordinate struct {
	Decimal    float32 `json:"decimal"`
	DMS        string  `json:"dms"`
	Degrees    int     `json:"degrees"`
	Minutes    int     `json:"minutes"`
	Seconds    float64 `json:"seconds"`
	Hemisphere rune    `json:"hemisphere"`
}

func (coor coordinate) inputDeg() (deg int) {
	fmt.Println("Please input the Degrees")
	_, err := fmt.Scanln(&deg)
	if err != nil {
		fmt.Println("Error encountered in coordinate.inputDeg")
		fmt.Println(err)
	}
	return deg
}

func (coor coordinate) inputMin() (min int) {
	fmt.Println("Please input the Minutes")
	_, err := fmt.Scanln(&min)
	if err != nil {
		fmt.Println("Error encountered in coordinate.inputMin")
		fmt.Println(err)
	}
	return min
}

func (coor coordinate) inputSec() (sec float64) {
	fmt.Println("Please input the Seconds")
	_, err := fmt.Scanln(&sec)
	if err != nil {
		fmt.Println("Error encountered in coordinate.inputSec")
		fmt.Println(err)
	}
	return sec
}

func (coor coordinate) inputhem() (hem rune) {
	fmt.Println("Please input the Hemisphere")
	_, err := fmt.Scanln(&hem)
	if err != nil {
		fmt.Println("Error encountered in coordinate.inputhem")
		fmt.Println(err)
	}
	unicode.ToUpper(hem)
	return hem
}

func (coor coordinate) generateDMS() (dms string) {
	dms = string(coor.Degrees) + "°"
	dms += string(coor.Minutes) + "'"
	dms += strconv.FormatFloat(coor.Seconds, 'E', -1, 32)
	dms += string(coor.Hemisphere)

	return dms
}

func main() {
	var location coordinate

	location.Degrees = location.inputDeg()
	location.Minutes = location.inputMin()
	location.Seconds = location.inputSec()
	location.Hemisphere = location.inputhem()

	location.DMS = location.generateDMS()

	fmt.Println(location.DMS)
}
