package main

import "fmt"

func main() {

	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	terrestrial := planets[0:4:4] //length 4 capacity 4
	worlds := append(terrestrial, "Ceres")

	fmt.Println(planets) //Prints [Mercury Venus Earth Mars Jupiter Saturn Uranus Neptune]

	terrestrial = planets[0:4] //length 4, capacity 8
	worlds = append(terrestrial, "Ceres")

	fmt.Println(planets) //Prints [Mercury Venus Earth Mars Ceres Saturn Uranus Neptune]

	fmt.Println(worlds)
}

/*It is better to use Three-Index Slicing unless you need to overwrite the elements of the underlying array. This
is because it is far safer to set the capacity with a Three-Index Slice. */
