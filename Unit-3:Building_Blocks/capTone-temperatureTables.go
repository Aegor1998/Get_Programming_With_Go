package main

import "fmt"

type (
	cel float64
	fah float64
)

func (c cel) Fahrenheit() fah {
	return fah((c + 32) * (9.0 / 5.0))
}

func (f fah) Celsius() cel {
	return cel((f - 32) * (5.0 / 9.0))
} //fah to cel

func main() {
	drawTableCelToFah()
	fmt.Printf("\n\n\n\n\n\n")
	drawTableFahToCel()
}

func drawTableCelToFah() {
	//Draw Table Header
	fmt.Println("===============")
	fmt.Println("|  °C  |  °F  |")
	fmt.Println("===============")
	//Draw Table Body
	for c := cel(-40.00); c <= 100.00; c += 5 {
		f := c.Fahrenheit()
		fmt.Printf("|%6.2f|%6.2f|\n", c, f)
	}
	//Draw Table End
	fmt.Println("===============")
}

func drawTableFahToCel() {
	//Draw Table Header
	fmt.Println("===============")
	fmt.Println("|  °F  |  °C  |")
	fmt.Println("===============")
	//Draw Table Body
	for f := fah(-40.00); f <= 100.00; f += 5 {
		c := f.Celsius()
		fmt.Printf("|%6.2f|%6.2f|\n", f, c)
	}
	//Draw Table End
	fmt.Println("===============")
}
