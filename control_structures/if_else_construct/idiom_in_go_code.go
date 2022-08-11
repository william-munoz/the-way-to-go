package main

import "fmt"

func main() {
	n := 78
	if n%2 != 0 {
		panic("need a even value")
	}
	fmt.Printf("Continue with an even value\n")
}
