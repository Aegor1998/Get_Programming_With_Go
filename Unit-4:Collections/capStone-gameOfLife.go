package main

import "fmt"

const (
	width  = 80
	height = 15
)

type universe [][]bool

func (u universe) Show() {
	fmt.Printf("|")
	for _, b := range u {
		switch b {
		case true:
			fmt.Printf(" * |")
		case false:
			fmt.Printf(" _ |")

		}
	}
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
