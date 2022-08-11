package main

import "fmt"

func main() {
	f1 := inc1

	v := f1(1)
	fmt.Printf("Value of v: %v\n", v)
}

func inc1(x int) int {
	return x + 1
}
