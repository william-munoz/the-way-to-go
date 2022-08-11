package main

import "fmt"

func main() {
	slice1 := make([]int, 4)
	fmt.Printf("slice1 len is: %d\n", len(slice1))
	// initialize the slice1 values
	for i := 0; i < len(slice1); i++ {
		slice1[i] = i + 1
	}

	for ix, value := range slice1 {
		fmt.Printf("Slice at %d is: %d\n", ix, value)
	}

	printSeasons()
}

func printSeasons() {
	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}
	for ix, season := range seasons {
		fmt.Printf("Season %d is: %s\n", ix, season)
	}

	var season string
	for _, season = range seasons {
		fmt.Printf("%s\n", season)
	}
}
