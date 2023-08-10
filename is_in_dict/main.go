package main

import (
	"fmt"
	"strings"
)

/*

function setup() {}
function isInDict() {}

idea is to to as much preparation in setup to optimize isInDict.

for example:

setup([]string{"car", "cat", "bar"})
isInDict("car") => true

It should also be possible to have a wildmark:

setup([]string{"car", "cat", "bar", "*ati"})
isInDict("pati") => true
*/

type Dictionary map[string]bool

func (dictionary Dictionary) isInDict(word string) bool {
	return dictionary[word]
}

func setup(words []string) (dictionary Dictionary) {
	dictionary = make(Dictionary)

	for _, word := range words {
		idx := strings.Index(word, "*")
		if idx != -1 {
			for c := 'a'; c <= 'z'; c++ {
				wordVariation := []rune(word)
				wordVariation[idx] = c
				dictionary[string(wordVariation)] = true
			}
		} else {
			dictionary[word] = true
		}
	}

	return
}

func main() {
	dictionary := setup([]string{"car", "*at", "bar"})

	fmt.Println(dictionary)
	fmt.Println(dictionary.isInDict("car"))  // true
	fmt.Println(dictionary.isInDict("cat"))  // true
	fmt.Println(dictionary.isInDict("bat"))  // true
	fmt.Println(dictionary.isInDict("subpar")) // false
}
