package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Code orginally said AD. Changed it to CE for Common Era
	const era = "CE"
	var (
		year  int
		month int
		days  int
		leap  bool
		day   int
	)

	year, month = getYearMonth()
	days = daysInAMonth(month)
	leap = isItALeapYear(year)
	days = addDayToFeb(days, leap)
	day = whatDay(days)
	monthName := time.Month(month)
	//Month Day, Year Era
	fmt.Printf("The Date is %v %v, %d %v\n", monthName, day, year, era)
	if leap == false {
		fmt.Println("This year is not a leap year")
	} else {
		fmt.Println("This year is a leap year")
	}
}

func daysInAMonth(month int) int {
	switch month {
	case 4, 6, 9, 11: //Checks for months with 31 days
		return 31
	case 2: //checks for february
		return 28
	default: //All other months have 30 days
		return 30
	}
}

func isItALeapYear(year int) bool {
	var leap bool = year%400 == 0 || (year%4 == 0 && year%100 != 0)
	return leap
}

func addDayToFeb(days int, leap bool) int {
	if days == 28 && leap == true {
		return days + 1
	}
	return days
}

func getYearMonth() (year int, month int) {
	//Seeding the rand function
	rand.Seed(time.Now().UnixNano())
	//randIntn() returns an int [0, MaxInt]
	year = rand.Intn(10000)
	month = rand.Intn(12) + 1
	return year, month
}

func whatDay(days int) int {
	//Seeding the rand function
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(days-1) + 1
}
