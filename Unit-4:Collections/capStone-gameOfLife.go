package main

import "fmt"

const (
	width  = 80
	height = 15
)

type universe [][]bool

func (u universe) Show() {
	fmt.Printf("|")
	for i, b := range u {
		switch b[i] {
		case true:
			fmt.Printf(" * |")
		case false:
			fmt.Printf(" _ |")

		}
	}
}

func (u universe) Seed() {
	/*use rand and time to populate universe
	1/4 of the cells are to be counted as alive
	*/
}

func main() {
	var (
		currentGeneration = newUniverse()
	)
	currentGeneration.Show()
}

func newUniverse() universe {
	var builtUniverse universe
	return builtUniverse
} //makes new universes
