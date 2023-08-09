package main

import "fmt"

/*
	An *even-ended* number is a number with the same
	first and last digits.

	Count how many even-ended numbers result from multiplying any two
	four-digit numbers.
*/

func isEvenEnded(n int) bool {
	nStr := fmt.Sprintf("%d", n)
	firstDigit := nStr[0]
	lastDigit := nStr[len(nStr)-1]

	return firstDigit == lastDigit
}

func main() {
	count := 0

	for i := 1000; i <= 9999; i++ {
		for j := i; j <= 9999; j++ {
			if isEvenEnded(i * j) {
				count++
			}
		}
	}

	// print the count
	fmt.Println(count)
}
