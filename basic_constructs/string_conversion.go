package main

import (
	"fmt"
	"strconv"
)

func main() {
	orig := "777"
	var an int
	var newS string
	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	an, _ = strconv.Atoi(orig) // converting to number
	fmt.Printf("The integer is: %d\n", an)
	an += 5
	newS = strconv.Itoa(an) // converting to string
	fmt.Printf("The new string is: %s\n", newS)
}
