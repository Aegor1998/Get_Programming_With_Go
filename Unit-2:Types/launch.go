package main

import "fmt"

func main() {
	var (
		launch     = false
		launchText = fmt.Sprintf("%v", launch)
		yesNo      string
	)
	fmt.Println("Ready for launch: ", launchText)

	if launch {
		yesNo = "yes"
	} else {
		yesNo = "no"
	}
	fmt.Println("Ready for launch: ", yesNo)
}
