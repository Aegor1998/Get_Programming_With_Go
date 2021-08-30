package main

import (
	"fmt"
	"log"
	"strings"
)

type temperature struct {
	name  string
	value float64
}

const (
	kelvin     string = "Kelvin"
	celsius    string = "Celsius"
	fahrenheit string = "Fahrenheit"
) //have this to make sure I spelt everything the same when comparing

func main() {
	var (
		convertFrom temperature
		ConvertTo   temperature
	)
	fillNameConvertFromAndTo(&convertFrom, &ConvertTo)
	getConvertFromValue(&convertFrom)
	convertFromTo(&convertFrom, &ConvertTo)
	displayConversion(&convertFrom, &ConvertTo)
}

//_________________________________________________________________________________________________________________
func fillNameConvertFromAndTo(convertFrom *temperature, ConvertTo *temperature) {
	var userChoice string
	instructionsForFillNameConvertFromAndTo()
	userChoice = getUserInputLower()
	convertFrom.name = fillTemperatureName(userChoice)
	if convertFrom.name == "ERROR" {
		fmt.Println("Error occurred in func fillTemperatureName while filling out convertFrom.name")
	} //Error catch for empty convertFrom.name
	instructionsForFillNameConvertFromAndTo()
	userChoice = getUserInputLower()
	ConvertTo.name = fillTemperatureName(userChoice)
	if convertFrom.name == "ERROR" {
		fmt.Println("Error occurred in func fillTemperatureName while filling out convertTo.name")
	} //Error catch for empty convertTo.name
}

//Get User Input for String for fillNameConvertFromAndTo
func getUserInputLower() (output string) {
	_, err := fmt.Scanln(&output)
	if err != nil {
		fmt.Println("Error Encountered: getUserInputLower() line 1")
		log.Fatal(err)
	}
	output = strings.ToLower(output)
	if output != "k" && output != "c" && output != "f" {
		fmt.Println("Incorrect Information entered. Please Try again")
		fmt.Print("New Entry: ")
		output = getUserInputLower()
	} //Incorrect Data Handling
	return output
}

//switch statement for fillNameConvertFromAndTo
func fillTemperatureName(userChoice string) (name string) {
	switch userChoice {
	case "k":
		return kelvin
	case "c":
		return celsius
	case "f":
		return fahrenheit
	default:
		return "ERROR"
	}

}

//fmt.print statements for fillNameConvertFromAndTo
func instructionsForFillNameConvertFromAndTo() {
	fmt.Println("What do you want to convert From?")
	fmt.Printf("For Kelvin enter k\nFor Celcius enter c\nFor Farenheit enter f\n")
	fmt.Print("User Choice: ")
}

//_________________________________________________________________________________________________________________
func getConvertFromValue(convertFrom *temperature) {
	fmt.Printf("What is the temperature measured in %s?\n", convertFrom.name)
	fmt.Print("User Choice: ")
	convertFrom.value = getUserInputFloat64()
	checkForLessThanAbsoluteZero(convertFrom)
}

//Get User Input for float64 for getConvertFromValue
func getUserInputFloat64() (output float64) {
	_, err := fmt.Scanln(&output)
	if err != nil {
		fmt.Println("Error Encountered: getUserInputFloat64() line 1")
		log.Fatal(err)
	}
	return output
}

//Deals with less than absolute zero for getConvertFromValue
func checkForLessThanAbsoluteZero(convertFrom *temperature) {
	const (
		absoluteZeroK float64 = 0
		absoluteZeroC float64 = -273.15
		absoluteZeroF float64 = -459.67
	)

	switch convertFrom.name {
	case kelvin:
		if convertFrom.value <= absoluteZeroK {
			fmt.Println("Out of range Value given: Temp cannot be lower than absolute zero")
			getConvertFromValue(convertFrom)
		}
	case celsius:
		if convertFrom.value <= absoluteZeroC {
			fmt.Println("Out of range Value given: Temp cannot be lower than absolute zero")
			getConvertFromValue(convertFrom)
		}
	case fahrenheit:
		if convertFrom.value <= absoluteZeroF {
			fmt.Println("Out of range Value given: Temp cannot be lower than absolute zero")
			getConvertFromValue(convertFrom)
		}
	}
}

//_________________________________________________________________________________________________________________
func convertFromTo(convertFrom *temperature, convertTo *temperature) {
	switch convertFrom.name {
	case kelvin:
		convertFromKelvin(convertFrom, convertTo)
	case celsius:
		convertFromCelsius(convertFrom, convertTo)
	case fahrenheit:
		convertFromFahrenheit(convertFrom, convertTo)
	}
	if convertTo.value == 0 {
		fmt.Println("Error Occurred: No Conversion saved")
	} //Error Message for if the conversion does not happen
}

//called in convertFromTo if convertFrom.name is equal to const kelvin
func convertFromKelvin(convertFrom *temperature, convertTo *temperature) {
	switch convertTo.name {
	case kelvin:
		convertTo.value = convertFrom.value
	case celsius:
		convertTo.value = convertFrom.value - 273.15
	case fahrenheit:
		convertTo.value = (convertFrom.value-273.15)*1.8 + 32
	}
}

//called in convertFromTo if convertFrom.name is equal to const celsius
func convertFromCelsius(convertFrom *temperature, convertTo *temperature) {
	switch convertTo.name {
	case kelvin:
		convertTo.value = convertFrom.value + 273.15
	case celsius:
		convertTo.value = convertFrom.value
	case fahrenheit:
		convertTo.value = (convertFrom.value * (9 / 5)) + 32
	}
}

//called in convertFromTo if convertFrom.name is equal to const fahrenheit
func convertFromFahrenheit(convertFrom *temperature, convertTo *temperature) {
	switch convertTo.name {
	case kelvin:
		convertTo.value = ((convertFrom.value - 32) * (5 * 9)) + 273.15
	case celsius:
		convertTo.value = (convertFrom.value - 32) * (5 * 9)
	case fahrenheit:
		convertTo.value = convertFrom.value
	}
}

//_________________________________________________________________________________________________________________
func displayConversion(convertFrom *temperature, convertTo *temperature) {
	fmt.Printf("%.2f %s >> %.2f %s", convertFrom.value, convertFrom.name, convertTo.value, convertTo.name)
}
