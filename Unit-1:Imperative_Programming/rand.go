package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	test()
	//randomDist()
}

func test() {
	rand.Seed(time.Now().UnixNano())
	var num = rand.Intn(2) + 1
	fmt.Println(num)

	num = rand.Intn(2) + 1
	fmt.Println(num)

	num = rand.Intn(2) + 1
	fmt.Println(num)

	num = rand.Intn(2) + 1
	fmt.Println(num)

	num = rand.Intn(2) + 1
	fmt.Println(num)

	num = rand.Intn(2) + 1
	fmt.Println(num)
}

func randomDist() {
	//for range (upperLimit - lowerLimit) + lowerLimit
	var disntance = rand.Intn(345000000) + 56000000
	fmt.Println(disntance)
}
