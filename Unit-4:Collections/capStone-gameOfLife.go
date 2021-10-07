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
	if u.Alive(x, y) != true {
		return 0
	}
	for column := -1; column < 2; column++ {
		switch {
		case (x + column) > 14:
			xCoordinate = 0 //This should only wrap around by a max of 1
		case (x + column) < 0:
			xCoordinate = 14 //This should only wrap around by a max of 1
		default:
			xCoordinate = x + column

		} //Keeps xCoordinate in range of 0 to 14
		for row := -1; row < 2; row++ {
			switch {
			case (y + row) > 79:
				yCoordinate = 0 //This should only wrap around by a max of 1
			case (y + row) < 0:
				yCoordinate = 79 //This should only wrap around by a max of 1
			default:
				yCoordinate = y + row
			} //Keeps yCoordinate in range of 0 to 79
			if u.Alive(xCoordinate, yCoordinate) == true {
				neighbors++
			}
		}
	}
	return neighbors
} //Counts the living neighbors and itself.

func main() {
	var (
		universeMap = make(map[int]universe) //1-d map containing a 2-d array
		//generations = numberOfGenerations()
	)
	universeMap[0] = universeMap[0].Seed()
	universeMap[0].Show()
	fmt.Printf("\nthe N count of 0,0: %v\n", universeMap[0].Neighbors(0, 0))
	fmt.Printf("the N count of 5,18: %v\n", universeMap[0].Neighbors(5, 18))
	fmt.Printf("the N count of 13,79: %v\n", universeMap[0].Neighbors(13, 79))

}

/*func numberOfGenerations() (gen int) {
	fmt.Print("How many generations will this last: ")
	fmt.Scanln(&gen)
	return gen
}*/
