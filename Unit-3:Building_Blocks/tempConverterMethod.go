package main

import (
	"fmt"
	"log"
	"strings"
)

type (
	kelvinM      float64
	celciusM     float64
	fahrenheitM  float64
	temperatureM struct {
		convertFrom string
		kelvin      kelvinM
		celcius     celciusM
		fahrenheit  fahrenheitM
	}
)

const (
	kelvinC     string = "Kelvin"
	celsiusC    string = "Celsius"
	fahrenheitC string = "Fahrenheit"
) //have this to make sure I spelt everything the same when comparing

//Methods: Convert From Kelvin
func (k kelvinM) celcius() celciusM {
	return celciusM(k - 273.15)
} //kelvinM >> celciusM
func (k kelvinM) fahrenheit() fahrenheitM {
	temp := k.celcius()
	return fahrenheitM((temp - 32) * (9.0 / 5.0))
} //kelvinM >> fahrenheitM

//Methods: Convert From Celcius
func (c celciusM) kelvin() kelvinM {
	return kelvinM(c + 273.15)
} //celciusM >> kelvinM
func (c celciusM) fahrenheit() fahrenheitM {
	return fahrenheitM((c + 32) * (9.0 / 5.0))
} //celciusM >> fahrenheitM

//Methods: Convert From Fahrenheit
func (f fahrenheitM) celcius() celciusM {
	return celciusM((f - 32) * (5.0 / 9.0))
} //fahrenheitM >> celciusM
func (f fahrenheitM) kelvin() kelvinM {
	temp := f.celcius()
	return kelvinM(temp + 273.15)
} //fahrenheitM >> kelvinM

func main() {
	var convertingStruct temperatureM

	fillNameConvertFrom(&convertingStruct)
	convertFromValue(&convertingStruct)
	fromTo(&convertingStruct)
	showConversion(&convertingStruct)
}

func fillNameConvertFrom(convertingStruct *temperatureM) {
	var userChoice string
	instructionsForFromAndTo()
	userChoice = userInputLower()
	convertingStruct.convertFrom = temperatureName(userChoice)
	if convertingStruct.convertFrom == "ERROR" {
		fmt.Println("Error occurred in func temperatureName while filling out convertFrom.name")
	} //Error catch for empty convertFrom.name
}

//Get User Input for String for fillNameConvertFromAndTo
func userInputLower() (output string) {
	_, err := fmt.Scanln(&output)
	if err != nil {
		fmt.Println("Error Encountered: userInputLower() line 1")
		log.Fatal(err)
	}
	output = strings.ToLower(output)
	if output != "k2" && output != "c" && output != "f" {
		fmt.Println("Incorrect Information entered. Please Try again")
		fmt.Print("New Entry: ")
		output = userInputLower()
	} //Incorrect Data Handling
	return output
}

//switch statement for fillNameConvertFromAndTo
func temperatureName(userChoice string) (name string) {
	switch userChoice {
	case "k2":
		return kelvinC
	case "c":
		return celsiusC
	case "f":
		return fahrenheitC
	default:
		return "ERROR"
	}

}

//fmt.print statements for fillNameConvertFromAndTo
func instructionsForFromAndTo() {
	fmt.Println("What do you want to convert From?")
	fmt.Printf("For Kelvin enter k2\nFor Celcius enter c\nFor Farenheit enter f\n")
	fmt.Print("User Choice: ")
}

//_________________________________________________________________________________________________________________
func convertFromValue(convertingStruct *temperatureM) {
	fmt.Printf("What is the temperature measured in %s?\n", convertingStruct.convertFrom)
	fmt.Print("User Choice: ")
	switch convertingStruct.convertFrom {
	case kelvinC:
		convertingStruct.kelvin = kelvinM(userInputFloat64())
	case celsiusC:
		convertingStruct.celcius = celciusM(userInputFloat64())
	case fahrenheitC:
		convertingStruct.fahrenheit = fahrenheitM(userInputFloat64())
	}
	absoluteZero(convertingStruct)
}

//Get User Input for float64 for convertFromValue
func userInputFloat64() (output float64) {
	_, err := fmt.Scanln(&output)
	if err != nil {
		fmt.Println("Error Encountered: userInputFloat64() line 1")
		log.Fatal(err)
	}
	return output
}

//Deals with less than absolute zero for convertFromValue
func absoluteZero(convertingStruct *temperatureM) {
	const (
		absoluteZeroK kelvinM     = 0
		absoluteZeroC celciusM    = -273.15
		absoluteZeroF fahrenheitM = -459.67
	)

	switch convertingStruct.convertFrom {
	case kelvinC:
		if convertingStruct.kelvin < absoluteZeroK {
			fmt.Println("Out of range Value given: Temp cannot be lower than absolute zero")
			convertFromValue(convertingStruct)
		}
	case celsiusC:
		if convertingStruct.celcius < absoluteZeroC {
			fmt.Println("Out of range Value given: Temp cannot be lower than absolute zero")
			convertFromValue(convertingStruct)
		}
	case fahrenheitC:
		if convertingStruct.fahrenheit < absoluteZeroF {
			fmt.Println("Out of range Value given: Temp cannot be lower than absolute zero")
			convertFromValue(convertingStruct)
		}
	}
}

//_________________________________________________________________________________________________________________
func fromTo(convertingStruct *temperatureM) {
	switch convertingStruct.convertFrom {
	case kelvinC:
		convertingStruct.celcius, convertingStruct.fahrenheit = convertingStruct.kelvin.celcius(), convertingStruct.kelvin.fahrenheit()
	case celsiusC:
		convertingStruct.kelvin, convertingStruct.fahrenheit = convertingStruct.celcius.kelvin(), convertingStruct.celcius.fahrenheit()
	case fahrenheitC:
		convertingStruct.kelvin, convertingStruct.celcius = convertingStruct.fahrenheit.kelvin(), convertingStruct.fahrenheit.celcius()
	}
}

//_________________________________________________________________________________________________________________
func showConversion(convertingStruct *temperatureM) {
	switch convertingStruct.convertFrom {
	case kelvinC:
		fmt.Printf("%.2f K == %.2f C == %.2f F", convertingStruct.kelvin, convertingStruct.celcius, convertingStruct.fahrenheit)
	case celsiusC:
		fmt.Printf("%.2f C == %.2f K == %.2f F", convertingStruct.celcius, convertingStruct.kelvin, convertingStruct.fahrenheit)
	case fahrenheitC:
		fmt.Printf("%.2f F == %.2f C == %.2f K", convertingStruct.fahrenheit, convertingStruct.celcius, convertingStruct.kelvin)
	} //Chooses in what order to print out values
}
