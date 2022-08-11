package main

import "fmt"

func main() {
	i1 := 5
	fmt.Printf("An integer: %d, and it's location in memory is: %p\n", i1, &i1)

	var intP *int
	intP = &i1

	fmt.Printf("intP variable point to: %p, and it's value is: %d\n", &intP, *intP)
}
