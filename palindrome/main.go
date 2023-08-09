package main

// Approach 1: Retrieve length.
// loop from 0 till the half-way point
// compare byte at index with length-index-1
func isPalindromeOne(inputString string) bool {
	length := len(inputString)

	for i := 0; i < length/2; i++ {
		if inputString[i] != inputString[length-i-1] {
			return false
		}
	}

	return true
}

// Approach 2: Define two pointers
// increment/decrement both on each for loop iteration
func isPalindromeTwo(inputString string) bool {
	a := 0
	b := len(inputString)-1

	for a <= b {
		if inputString[a] != inputString[b] {
			return false
		}

		a++
		b--
	}

	return true
}
