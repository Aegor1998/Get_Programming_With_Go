/*
Program Goals:
	1) list of 3 animals
	2) Each animal should have a name and use the Stringer Interface to return their name
	3) Every animal should have a move and eat method
		a) The move method returns a desription of movement
		b) The eat method returns a random food as that animal's favorite
	4) Implement a 24hr day/night cycle (12 Day + 12 Night)
	5) Run the simulation for 3 day cycles
		a) Simulation
			1) Every hour of the day pick an animal at random and perform a random action (move or eat)
			2) For every action print out a description of what the animal did
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type animal struct {
	name string
	eat  string
	move string
}

//methods
func (a *animal) sayer(what string) {
	switch what {
	case "name":
		fmt.Printf("\nThe animal's name is %s", a.name)
	case "eat":
		fmt.Printf("\nThe animal's food is %s", a.eat)
	case "move":
		fmt.Printf("\nThe animal's movement is %s", a.move)
	}
}

func (a *animal) feeder() {
	a.eat = FOOD[randomNumber(len(FOOD))]
}

func (a *animal) mover() {
	a.move = MOVEMENT[randomNumber(len(MOVEMENT))]
}

//static arrays
var ANIMAL = [3]string{"Dog", "Gopher", "Rabbit"}
var FOOD = [4]string{"hay", "apples", "roots", "chow"}
var MOVEMENT = [3]string{"bipedal", "Quadrupedal", "Swim"}

func main() {
	var (
		simulationSet []animal
		cycles        int
	)
	//Load Array
	for i := 1; i < len(ANIMAL); i++ {
		var ani animal
		ani.name = ANIMAL[i]
		simulationSet = append(simulationSet, animal{name: ani.name})
	} //not in its own function because the IDE complained,
	// and I could figure out how to fix it
	cycles = 3
	runSimulation(cycles, simulationSet)
}

func runSimulation(cycles int, testSet []animal) {

	for c := 0; c < cycles; c++ {
		fmt.Println("Start of Cycle ", c+1)
		for h := 0; h < 24; h++ {
			switch {
			case h >= 12:
				//do nothing
			default:
				var subject animal
				subject.name = testSet[randomNumber(len(testSet))].name
				subject.feeder()
				subject.mover()
				subject.sayer("name")
				subject.sayer("eat")
				subject.sayer("move")
			}
		}
	}
}

func randomNumber(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - 1)
}
