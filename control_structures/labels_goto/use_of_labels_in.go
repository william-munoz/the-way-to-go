package main

import "fmt"

func main() {
LABEL1: // adding a label for location
	for i := 0; i <= 5; i++ { // outer loop
		for j := 0; j <= 5; j++ { // inner loop
			if j == 4 {
				continue LABEL1 // jump to the label LABEL1
			}
			fmt.Printf("i is %d, and j is %d\n", i, j)
		}
	}
}
