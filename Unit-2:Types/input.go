package main

import "fmt"

func main() {
	test := [6]string{"true", "yes", "1", "false", "no", "0"}

	for i := 0; i < len(test); i++ {
		temp := convert(test[i])
		fmt.Printf("input: %s | Output: %v\n", test[i], temp)
	}
}

func convert(str string) (answer bool) {
	switch str {
	case "true", "yes", "1":
		answer = true
	case "false", "no", "0":
		answer = false
	}
	return answer
}
