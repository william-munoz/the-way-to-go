package main

import "fmt"

func main() {
	a := [15]int{}
	fmt.Printf("initial a content: %v\n", a)

	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	fmt.Printf("new a content: %v\n", a)

	for pos, item := range a {
		fmt.Printf("Position: %d - Value: %d\n", pos, item)
	}
}
