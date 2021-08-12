package main

import "fmt"

func main() {
	var (
		bh float64 = 32768
		h          = int16(bh)
	)
	fmt.Println(h)
}
