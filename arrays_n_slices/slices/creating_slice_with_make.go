package main

import "fmt"

func main() {
	s2 := make([]int, 10)
	fmt.Printf("values of s2 slice: %v\n", s2)

	fmt.Printf("capacity of s2 slice: %v\n", cap(s2))

	// load the array/slice:
	for i := 0; i < len(s2); i++ {
		s2[i] = 5 * i
	}

	// print the slice:
	for i := 0; i < len(s2); i++ {
		fmt.Printf("Slice at %d is %d\n", i, s2[i])
	}
	fmt.Printf("\nThe lenght of s2 is %d\n", len(s2))
	fmt.Printf("\nThe capacity of s2 is %d\n", cap(s2))
}
