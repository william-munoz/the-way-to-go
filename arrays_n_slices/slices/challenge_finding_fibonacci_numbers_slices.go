package main

import "fmt"

func main() {
	fmt.Printf("fibarray(2): %v\n", fibarray(2))
	fmt.Printf("fibarray(4): %v\n", fibarray(4))
	fmt.Printf("fibarray(5): %v\n", fibarray(5))
	fmt.Printf("fibarray(10): %v\n", fibarray(10))

	fmt.Printf("fibarray2(10): %v\n", fibarray2(10))
}

// my solution
func fibarray(term int) []int {
	s := make([]int, term)
	for i := 0; i < len(s); i++ {
		if i == 0 {
			s[i] = 0
		} else if i == 1 {
			s[i] = 1
		} else {
			s[i] = s[i-2] + s[i-1]
		}
	}
	return s
}

// author solution
func fibarray2(term int) []int { // calculating Fibonacci for first term numbers
	farr := make([]int, term)

	//base case
	if term == 1 {
		farr[0] = 0
		return farr
	}
	farr[0], farr[1] = 0, 1

	// recursive case
	for i := 2; i < term; i++ {
		farr[i] = farr[i-1] + farr[i-2] // calculating sequence for i index
	}
	return farr
}
