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
	drawTableHeader("C", "F")
	//Draw Table Body
	drawTableEntry("C")
	//Draw Table End
	fmt.Println("===============")
}

func drawTableFahToCel() {
	//Draw Table Header
	drawTableHeader("F", "C")
	//Draw Table Body
	drawTableEntry("F")
	//Draw Table End
	fmt.Println("===============")
}

func drawTableHeader(pos1 string, pos2 string) {
	fmt.Println("===============")
	fmt.Printf("|  °%s  |  °%s  |\n", pos1, pos2)
	fmt.Println("===============")
}

func drawTableEntry(kind string) {
	switch kind {
	case "C":
		for c := cel(-40.00); c <= 100.00; c += 5 {
			fmt.Printf("|%6.2f|%6.2f|\n", c, c.Fahrenheit())
		}
	case "F":
		for f := fah(-40.00); f <= 100.00; f += 5 {
			fmt.Printf("|%6.2f|%6.2f|\n", f, f.Celsius())
		}
	}
}
