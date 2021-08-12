package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"text/tabwriter"
	"time"
)

//rand.Seed(time.Now().UnixNano())
type flight struct {
	companyName string
	travelTime  int    //days
	tripType    string //Round-Trip or One-Way
	price       int    //in millions of dollars
}

const (
	DEPMON  string = string(time.Month(10))
	DEPDAY  int    = 13
	DEPYEAR int    = 2020
)

func main() {
	//tripInfo := make([]flight, 0, 10)
	//Initilize tabwriter
	writer := new(tabwriter.Writer)
	// minwidth, tabwidth, padding, padchar, flags
	writer.Init(os.Stdout, 0, 8, 2, '\t', 0)
	defer writer.Flush()

	//Top of Table
	fmt.Fprintf(writer, "\n %s\t%s\t%s\t%s\t", "SpaceLine", "Days", "Trip", "Price")
	fmt.Fprintf(writer, "\n %s\t%s\t%s\t%s\t", "----------------", "----", "----------", "-----")

	for counter := 0; counter < 10; counter++ {
		var currentTrip flight //keeps the variable ina limited scope
		//adds values to the variable
		currentTrip.companyName = getCompnaieName()
		currentTrip.travelTime = getTravelTime()
		currentTrip.tripType = getTripType()
		currentTrip.price = getPrice(currentTrip.tripType)
		if currentTrip.companyName == "SpaceX" {
			currentTrip.companyName = currentTrip.companyName + "         "
		}
		fmt.Fprintf(writer, "\n %s\t%d\t%s\t%d\t\n", currentTrip.companyName, currentTrip.travelTime, currentTrip.tripType, currentTrip.price)
	}
}

func getCompnaieName() string {
	COMPANIES := [3]string{0: "Space Adventures", 1: "SpaceX", 2: "Virgin Galactic"}
	rand.Seed(time.Now().UnixNano())
	return COMPANIES[rand.Intn(3)]
}

func getTravelTime() int {
	rand.Seed(time.Now().UnixNano())
	const (
		DISTMARS int = 62100000 // km
		MAXSPEED int = 2592000  // 43,200 km/d
		MINSPEED int = 1382400  // 23,040 km/d
	)
	//Round to the nearest whole number(distance/speed) = travelTime
	return int(math.Round(float64(DISTMARS) / float64(rand.Intn(MAXSPEED-MINSPEED)+MINSPEED)))
}

func getTripType() string {
	rand.Seed(time.Now().UnixNano())
	var tripType int = rand.Intn(2) + 1
	if tripType == 1 {
		return "One-Way"
	}
	return "Round-Trip"
}

func getPrice(tripType string) int {
	rand.Seed(time.Now().UnixNano())
	const (
		MAXPRICE int = 50 // Million Dollars
		MINPRICE int = 36 // Million Dollars
	)
	var price int = rand.Intn(MAXPRICE-MINPRICE) + MINPRICE
	if tripType == "One-Way" {
		return price
	}
	return price * 2
}
