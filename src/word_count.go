package main 

import (
	"fmt"
	"strings"
)

func getWordCount(text string) map[string] int {

	wordMap := map[string]int {}

	words := strings.Fields(text)

	for _, word := range words {
		var loweredWord string = strings.ToLower(word)
		value, ok := wordMap[loweredWord]
		
		if ok {
			wordMap[loweredWord] = value +1
		} else {
			wordMap[loweredWord] = 1
		}
	}

	return wordMap
}

func main() {

	var text string = ` Needles and pins Needles and pins Sew me a Sail To catch me the wind`
	wordCountMap := getWordCount(text)

	fmt.Println(wordCountMap)
}