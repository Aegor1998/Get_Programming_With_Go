package main

import "fmt"

type turtle struct {
	xCoor int
	yCoor int
}

func (t *turtle) moveLeft() {
	t.xCoor = -1
}
func (t *turtle) moveRight() {
	t.xCoor = +1
}
func (t *turtle) moveUp() {
	t.yCoor = +1
}
func (t *turtle) moveDown() {
	t.yCoor = -1
}

func main() {
	var bob = turtle{xCoor: 0, yCoor: 0}

	// move bob 1 space up and 1 space right (1,1)
	bob.moveUp()
	bob.moveRight()
	fmt.Printf("\n(x,y) , (%d,%d)\n", bob.xCoor, bob.yCoor)

	//move bob 2 spaces down (1,-1)
	bob.moveDown()
	bob.moveDown()
	fmt.Printf("\n(x,y) , (%d,%d)\n", bob.xCoor, bob.yCoor)

	//move bob 2 spaces left (-1,-1)
	bob.moveLeft()
	bob.moveLeft()
	fmt.Printf("\n(x,y) , (%d,%d)\n", bob.xCoor, bob.yCoor)

	//move bob 2 spaces up (-1,1)
	bob.moveUp()
	bob.moveUp()
	fmt.Printf("\n(x,y) , (%d,%d)\n", bob.xCoor, bob.yCoor)
}
