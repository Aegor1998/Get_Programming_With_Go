package main

import "fmt"

type (
	kelvinT  float64
	celsiusT float64
)

func main() {
	var (
		k kelvinT
		c celsiusT
	)
	k = 294.0
	c = kelvinToCelcius(k)
	fmt.Printf("%.2f K is %.2f C\n", k, c)

	c = 127.0
	k = celciusToKelvin(c)
	fmt.Printf("%.2f C is %.2f K\n", c, k)
}

func kelvinToCelcius(k kelvinT) celsiusT {
	return celsiusT(k - 273.15)
}

func celciusToKelvin(c celsiusT) kelvinT {
	return kelvinT(c + 273.15)
}
