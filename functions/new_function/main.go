package main

import "fmt"

func main() {
	v := new(int) // creates a pointer to a T type
	fmt.Printf("v is a pointer: %v\n", v)
	*v = 10
	fmt.Printf("New value of v: %v\n", *v)
}
