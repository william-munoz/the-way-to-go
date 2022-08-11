package main

import (
	"fmt"
	"strconv"
)

func main() {
	orig := "ABC"
	var an int
	var err error
	// storing integer and error information
	an, err = strconv.Atoi(orig)

	if err != nil {
		fmt.Printf("is not possible to convert 'orig' variable with value '%s' to an integer - exiting with error\n", orig)
		return
	}

	fmt.Printf("The result integer is: %d\n", an)
}
