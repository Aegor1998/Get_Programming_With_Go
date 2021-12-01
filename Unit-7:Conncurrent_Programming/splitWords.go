package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		pipeline = make(chan string)
		sentence = "The dog barks at the cat"
	)

	go separations(pipeline, sentence)
	printPipeline(pipeline)
}

func separations(pipeline chan string, sentence string) {
	var (
		wordList []string
	)
	wordList = strings.Fields(sentence)
	for _, i := range wordList {
		pipeline <- i
	}
	close(pipeline)
}

func printPipeline(pipeline chan string) {
	for i := range pipeline {
		fmt.Println(i)
	}
}
