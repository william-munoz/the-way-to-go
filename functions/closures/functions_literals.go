package main

import "fmt"

func main() {
	psum := func() int {
		sum := 0
		for i := 0; i <= 1e6; i++ {
			sum += i
		}
		return sum
	}

	fmt.Printf("Sum of 1 million: %d", psum())
}
