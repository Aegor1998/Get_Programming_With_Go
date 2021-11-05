package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 15
	ALIVE  = "*" //Symbol used when printing out the map
	DEAD   = "_" //symbol used when printing out the map
)

type universe [height][width]bool

func (u universe) Show() {
	//row := make(map[int]bool, width)//This is to address each row of the map. Pointers have not been covered yet

	for h := 0; h < height; h++ {
		fmt.Printf("\nRow_%v	|", h) //Mostly here to start a newline for each row
		for w := 0; w < width; w++ {
			switch u[h][w] {
			case true:
				fmt.Printf("%s|", ALIVE)
			case false:
				fmt.Printf("%s|", DEAD)
			default:
				fmt.Println("Error Encountered in universe.show: No Value In Cell")
			} //I am using a const here for the dead/alive symbol so that I can easily change it
		}
	}
	fmt.Printf("\n        |")
	for i := 0; i < 8; i++ {
		for n := 0; n < 10; n++ {
			fmt.Printf("%v|", n)
		}
	}
}

func (u universe) Seed() universe {
	/*
		use rand and time to populate universe
		1/4 of the cells are to be counted as alive
	*/
	rand.Seed(time.Now().UnixNano())
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			chance := rand.Intn(4)
			switch chance { //using switch for readability
			case 2:
				u[h][w] = true
			default:
				u[h][w] = false
			}
		}
	}

	return u
}

func (u universe) Alive(x, y int) (state bool) {
	if x == 14 {
		x = 1
	}
	if x < 0 || y < 0 {
		return false
	} //Deals with negative coordinates
	if u[x][y] == true {
		state = true
	} else {
		state = false
	}
	return state
} //returns true if cell is alive

func (u universe) Neighbors(x, y int) (neighbors int) {
	var (
		xCoordinate int
		yCoordinate int
	)
	for row := -1; row < 2; row++ {
		switch {
		case (x + row) > 14:
			xCoordinate = 0 //This should only wrap around by a max of 1
		case (x + row) < 0:
			xCoordinate = 14 //This should only wrap around by a max of 1
		default:
			xCoordinate = x + row

		} //Keeps xCoordinate in range of 0 to 14
		for column := -1; column < 2; column++ {
			switch {
			case (y + column) > 79:
				yCoordinate = 0 //This should only wrap around by a max of 1
			case (y + column) < 0:
				yCoordinate = 79 //This should only wrap around by a max of 1
			default:
				yCoordinate = y + column
			} //Keeps yCoordinate in range of 0 to 79
			if u.Alive(xCoordinate, yCoordinate) == true {
				neighbors++
			}
		}
	}
	return neighbors
} //Counts the living neighbors and itself.

func (u universe) Next(x, y int) (living bool) {
	var (
		neighbor = u.Neighbors(x, y)
	)
	if u.Alive(x, y) == false {
		switch neighbor {
		case 3:
			living = true
		default:
			living = false
		}
	} else {
		switch neighbor {
		case 3: //was 2
			living = true
		case 4: //was 3
			living = true
		default:
			living = true
		} //had to add +1 to these because neighbors will report itself if alive
	}
	return living
}

func main() {
	var (
		universeMap = make(map[int]universe) //1-d map containing a 2-d array
		generations = numberOfGenerations()
		currentUni  = universeMap[0]
	)

	currentUni = universeMap[0].Seed()
	currentUni.Show()

	for i := 1; i < generations; i++ {
		nextUni := universeMap[i]

		nextGeneration(&currentUni, &nextUni)
		fmt.Printf("\n")
		nextUni.Show()
		currentUni = nextUni
	}
}

func numberOfGenerations() (gen int) {
	fmt.Print("How many generations will this last: ")
	fmt.Scanln(&gen)
	return gen
}

func nextGeneration(currentUni *universe, nextUni *universe) {
	for row := 0; row < height; row++ {
		for column := 0; column < width; column++ {
			nextUni[row][column] = currentUni.Next(row, column)
		}
	}
}
