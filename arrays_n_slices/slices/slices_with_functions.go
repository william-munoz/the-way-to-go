package main

import "fmt"

func main() {
	var arr = [5]int{0, 1, 2, 3, 4} // declare an array
	fmt.Println(sum(arr[:]))        // passing slice to the function
}

func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}
