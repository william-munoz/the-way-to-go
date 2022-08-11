package main

import "fmt"

func main() {
	fmt.Printf("Sum of list: %d\n", sumInts())
	fmt.Printf("Sum of list: %d\n", sumInts(7, 7, 7))
	fmt.Printf("Sum of list: %d\n", sumInts(3, 7, 17, 23, 27))
}

func sumInts(list ...int) (sum int) {
	for _, v := range list {
		sum += v
	}
	return // return the named variables, in this case: sum
}
