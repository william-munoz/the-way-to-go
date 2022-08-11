package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Printf("Value of i: %d\n", i)
		for j := 0; j < 10; j++ {
			if j > 5 {
				fmt.Printf("Breaking the inner loop on i: %d and j: %d\n", i, j)
				fmt.Println()
				break
			}
			fmt.Printf("Value of i: %d and j: %d\n", i, j)
		}
	}
}
