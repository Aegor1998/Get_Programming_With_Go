package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		text = "As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he felt able he ran again; the ground continued soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear except the fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man."
	)
	makeArrayOfWords(text)
}

func makeArrayOfWords(text string) {
	var (
		wordSlice = strings.Fields(strings.ToLower(text))
		wordCount = map[string]int{}
	)
	for _, i := range wordSlice {
		i = removeSymbols(i)

		_, found := wordCount[i]
		if found {
			wordCount[i] += 1
		} else {
			wordCount[i] = 1
		}
	} //Indexes each word as a key and records the number of occurrences.
	fmt.Println(wordCount)
}

func removeSymbols(word string) string {
	for _, i := range word {
		switch i {
		case '.', '!', ',', '?', '-', ';', ':':
			word = strings.Trim(word, string(i))
		}
	}
	return word
} //Removes symbols using strings.Trim()
