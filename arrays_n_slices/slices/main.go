package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	fmt.Printf("values on s slice: %v\n", s)

	// a slice made out of a slice
	s2 := s[:]
	fmt.Printf("values on s2 slice: %v\n", s2)

	// change on s slice
	s[0] = 7
	fmt.Printf("updated values on s slice: %v\n", s)

	// s2 points to s slice
	fmt.Printf("values on s2 slice: %v\n", s2)

	// append a value to s slice
	s = append(s, 7)
	fmt.Printf("updated values on s slice: %v\n", s)

	// new value on s slice is not reflected on s2 slice
	fmt.Printf("values on s2 slice: %v\n", s2)
}
