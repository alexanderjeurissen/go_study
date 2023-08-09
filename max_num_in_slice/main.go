package main

import "fmt"

func main() {
	nums := []int{16,8,42,4,23,15}

	maxVal := nums[0]

	for _,n := range nums[1:] {
		if n > maxVal {
			maxVal = n
		}
	}

	fmt.Println(maxVal)
}
