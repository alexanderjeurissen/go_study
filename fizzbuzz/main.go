package main

import "fmt"

/*
	FizBuzz

	Print numbers 1..20, however:
	-	If number divisible by 3 => print "fizz"
	-	If number divisible by 5 => print "buzz"
	-	If number divisible by both 3 and 5 => print "Fizz Buzz"
	-	Otherwise print number
*/

func main() {
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
