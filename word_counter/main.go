package main

import (
	"fmt"
	"strings"
)

/*
	Count how many times a word appears in the text.
*/

func wordCounts(str string) map[string]int {
	words := strings.Fields(str)

	counts := make(map[string]int)
	for _, word := range words {
		counts[word] += 1
	}

	return counts
}

func main() {
	str := "this is a string is it not"

	counts := wordCounts(str)

	fmt.Println(counts)
}
