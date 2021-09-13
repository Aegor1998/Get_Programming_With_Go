package main

import "fmt"

type stopCondition string //Not using rune because of issues with boolean expressions

func (r stopCondition) input() stopCondition {
	fmt.Scan(&r)
	return r
}

func main() {
	var (
		list                      = make([]int, 0, 10)
		stop        stopCondition = "n"
		currentSize int           = cap(list)
		oldSize     int
	)

	for i := 0; stop == "n"; i++ {
		oldSize = currentSize
		list = append(list, i)
		currentSize = len(list)

		if currentSize > oldSize {
			fmt.Printf("\nThe old capacity of list    : %d\n", oldSize)
			fmt.Printf("The current capacity of list: %d\n", currentSize)
			fmt.Printf("The contents of list    : %v\n", list)
			stop = stop.input()
		} else {
			fmt.Printf("The contents of list    : %v\n", list)
		}
	}
}
