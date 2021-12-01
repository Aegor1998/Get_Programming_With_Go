package main

import "fmt"

func main() {
	var (
		currentWord = make(chan string)
		wordList    = []string{"dog", "puppy", "road", "road", "cat", "mouse", "mouse", "cat"}
	)
	go ingestAndCheck(currentWord, wordList)
	output(currentWord)
}

func ingestAndCheck(currentWord chan string, wordList []string) {
	var lastWord string
	for _, i := range wordList {
		if i != lastWord {
			currentWord <- i
		}
		lastWord = i
	}
	close(currentWord)
}

func output(currentWord chan string) {
	fmt.Printf("\nBegin Printing: ")
	for i := range currentWord {
		fmt.Printf("%s ", i)
	}
}
