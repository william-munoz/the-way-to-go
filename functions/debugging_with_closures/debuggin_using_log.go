package main

import (
	"fmt"
	"log"
)

func main() {
	var where = log.Print
	where()
	// some code
	a := 2 * 5
	where()
	// some more code
	b := a + 100
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	where()
}
