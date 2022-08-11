package main

import "fmt"

func main() {
	i := 0
	// condition controlled for loop with 5 iterations
	for i < 5 {
		fmt.Printf("This is the %d iteration\n", i)
		i++
	}
}
