package main

import "fmt"

func main() {
	value := 0            // used to set values to 2D array
	screen := [2][2]int{} // will contain 4 values

	// assigning values in 2D-array
	for row := range screen {
		for column := range screen[row] {
			screen[row][column] = value
			value = value + 1
		}
	}

	// printing 2D-array
	for row := range screen {
		for column := range screen[row] {
			fmt.Print(screen[row][column], " ")
		}
		fmt.Printf("\n") // separates rows
	}
}
