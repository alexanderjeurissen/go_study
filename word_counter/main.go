package main

import (
	"fmt"
	"strings"
)

/*
	Count how many times a word appears in the text.
*/

func main() {
	// Use the Fields method in the String package  to split string in words
	// ToLower

	str := "this is a string"
	words := strings.Fields(str)
	fmt.Println()
}
