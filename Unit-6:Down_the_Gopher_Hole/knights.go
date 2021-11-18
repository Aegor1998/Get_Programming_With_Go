package main

import "fmt"

type character struct {
	name      string
	leftHand  item
	rightHand item
}

type item struct {
	name string
}

func (i *item) pickup(act *character, side string) {
	fmt.Printf("\n%v picks up %v\n", act.name, i.name)
	switch side {
	case "right":
		act.rightHand = *i
	case "left":
		act.leftHand = *i
	}
}

func (i *item) give(pOne, pTwo *character, side string) {
	fmt.Printf("\n%v gives %v to %v", pOne.name, i.name, pTwo.name)
	switch side {
	case "right":
		pTwo.rightHand = pOne.rightHand
		pOne.rightHand.name = ""
	case "left":
		pTwo.leftHand = pOne.leftHand
		pOne.leftHand.name = ""
	}
}

func main() {
	var (
		arthur = character{name: "Arthur"} //leftHand and rightHand kept empty
		knight = character{name: "Knight"} //leftHand and rightHand kept empty
		shield = item{name: "shield"}
		sword  = item{name: "sword"}
	)
	shield.pickup(&arthur, "left")
	shield.give(&arthur, &knight, "left")
	sword.pickup(&arthur, "right")
	sword.give(&arthur, &knight, "right")
}
