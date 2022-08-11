package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("This is the %d iteration.\n", i)
	}

	fmt.Println("---")

	for i := 10; i > 0; i-- {
		fmt.Printf("This is the %d iteration.\n", i)
	}
}
