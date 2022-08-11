package main

import (
	"fmt"
	"strings"
)

func main() {
	orig := "Hey, how are you George?"
	fmt.Printf("The original string is: %s\n", orig)
	lower := strings.ToLower(orig) // changing to lower case
	fmt.Printf("The lower case string is: %s\n", lower)
	upper := strings.ToUpper(orig) // changing to upper case
	fmt.Printf("The upper case string is: %s\n", upper)
}
