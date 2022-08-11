package main

import (
	"fmt"
	"strings"
)

func main() {
	seasons := []string{"spring", "summer", "autumn", "winter"}
	for ix, season := range seasons {
		fmt.Printf("season %d is: %s\n", ix, season)
	}
	for ix := range seasons {
		seasons[ix] = strings.Title(seasons[ix])
	}
	for ix, season := range seasons {
		fmt.Printf("season %d is: %s\n", ix, season)
	}
}
