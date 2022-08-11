package main

import "fmt"

func main() {
	n := 42
	if n%2 == 0 {
		fmt.Printf("The value is even\n")
	} else {
		fmt.Printf("The value is odd\n")
	}
}
