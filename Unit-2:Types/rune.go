package main

import "fmt"

func main() {
	var (
		pi    rune = 960
		alpha rune = 940
		omega rune = 969
		bang  byte = 33
	)
	fmt.Printf("%v, %v, %v, %v\n", pi, alpha, omega, bang)
	fmt.Printf("%c, %c, %c, %c\n", pi, alpha, omega, bang)
}
