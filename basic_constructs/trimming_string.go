package main

import (
	"fmt"
	"strings"
)

func main() {
	str := " This is an example of the string "
	fmt.Printf("Original string: %s\n", str)
	fmt.Printf("TrimSpace string: %s\n", strings.TrimSpace(str))

	trimSpecificString()
}

func trimSpecificString() {
	str := " This is an example of the string "
	fmt.Printf("Original String: %s\n", str)
	fmt.Printf("str's lenght: %d\n", len(str))
	stringToTrim := "string "
	fmt.Printf("Trim a specific string from str: %s\n", strings.Trim(str, stringToTrim))
	fmt.Printf("Result lenght: %d\n", len(strings.Trim(str, stringToTrim)))
}
