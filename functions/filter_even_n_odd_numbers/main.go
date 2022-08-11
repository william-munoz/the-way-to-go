package main

import (
	"fmt"
)

type flt func(int) bool // aliasing type

func main() {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Printf("slice value: %v\n", slice)
	even, odd := filter(slice, isEven)
	fmt.Printf("Even values: %v\n", even)
	fmt.Printf("Odd values: %v\n", odd)
}

func isEven(n int) bool { // check if n is even or not
	if n%2 == 0 {
		return true
	}
	return false
}

func filter(sl []int, f flt) (yes, no []int) { // split s into two slices: even and odd
	for _, val := range sl {
		if val%2 == 0 {
			yes = append(yes, val)
		} else {
			no = append(no, val)
		}
	}
	return
}
