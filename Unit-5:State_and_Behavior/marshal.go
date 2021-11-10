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
	Decimal    float64 `json:"decimal"`
	DMS        string  `json:"dms"`
	Degrees    float64 `json:"degrees"`
	Minutes    float64 `json:"minutes"`
	Seconds    float64 `json:"seconds"`
	Hemisphere rune    `json:"hemisphere"`
}

func (coor *coordinate) inputDeg() {
	fmt.Println("Please input the Degrees")
	_, err := fmt.Scanln(&coor.Degrees)
	if err != nil {
		fmt.Println("Error encountered in coordinate.inputDeg")
		fmt.Println(err)
	}
}

func (coor *coordinate) inputMin() {
	fmt.Println("Please input the Minutes")
	_, err := fmt.Scanln(&coor.Minutes)
	if err != nil {
		fmt.Println("Error encountered in coordinate.inputMin")
		fmt.Println(err)
	}

}

func (coor *coordinate) inputSec() {
	fmt.Println("Please input the Seconds")
	_, err := fmt.Scanln(&coor.Seconds)
	if err != nil {
		fmt.Println("Error encountered in coordinate.inputSec")
		fmt.Println(err)
	}
}

func (coor *coordinate) inputhem() {
	fmt.Println("Please input the Hemisphere")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error encountered in coordinate.inputhem")
		fmt.Println(err)
	}
	coor.Hemisphere = rune(input[0])
	unicode.ToUpper(coor.Hemisphere)
}

func (coor *coordinate) generateDMS() {
	coor.DMS = strconv.FormatFloat(coor.Degrees, 'f', 0, 32) + "° " + strconv.FormatFloat(coor.Minutes, 'f', 0, 32) + "' " + strconv.FormatFloat(coor.Seconds, 'f', 2, 32) + string(coor.Hemisphere)
}

func (coor *coordinate) generateDD() {
	coor.Decimal = coor.Degrees + ((coor.Minutes / 60) + (coor.Seconds / 3600))
	if coor.Hemisphere == 'S' {
		coor.Decimal = coor.Decimal * (-1)
	}
}

func main() {
	var location coordinate

	location.inputDeg()
	location.inputMin()
	location.inputSec()
	location.inputhem()

	location.generateDMS()
	location.generateDD()

	fmt.Println(location.DMS)
	fmt.Println(location.Decimal)
}
