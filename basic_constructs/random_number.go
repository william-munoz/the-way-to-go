package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNumber()
}

func randomNumber() {
	a := rand.Int()
	b := rand.Intn(8) // Random values of range: [0, n) (starting from 0 to n-1)

	fmt.Printf("a is: %d\n", a)
	fmt.Printf("b is: %d\n", b)
}
