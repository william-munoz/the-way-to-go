package main

import "fmt"

func main() {
	v := make([]int, 10, 10) // creates an empty array
	fmt.Printf("Value of v: %v\n", v)
	v[0] = 7
	fmt.Printf("Value of v[0]: %v\n", v[0])
	fmt.Printf("Value of v[1]: %v\n", v[1])
}
