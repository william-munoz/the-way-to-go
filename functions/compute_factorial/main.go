package main

import "fmt"

func main() {
	for i := uint64(0); i < uint64(22); i++ {
		fmt.Printf("Factorial of %d is %d\n", i, Factorial(i)) // calculating factorial of first 21 integers
	}
}

func Factorial(n uint64) (fac uint64) {
	if n <= 1 { // base case
		return 1
	}
	fac = n * Factorial(n-1) // recursive case
	return
}
