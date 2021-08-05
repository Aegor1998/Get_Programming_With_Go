package main

import "fmt"

func main() {
	var count uint8 = 0
	count-- //wraps back to 255
	fmt.Println(count)
}
