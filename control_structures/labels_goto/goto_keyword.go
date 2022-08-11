package main

import "fmt"

func main() {
	i := 0
HERE: // adding a label for a location
	fmt.Printf("%d", i)
	i++
	if i == 5 {
		return // get out from the main function
	}
	goto HERE // go to HERE label
}
