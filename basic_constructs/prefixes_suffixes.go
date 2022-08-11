package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "This is an example of a string"
	fmt.Printf("T/F? \nDoes the string \"%s\" have prefix %s? ", str, "Th")
	fmt.Printf("\n%t\n\n", strings.HasPrefix(str, "Th")) // Finding prefix

	fmt.Printf("Does the string \"%s\" have suffix %s? ", str, "ting")
	fmt.Printf("\n%t\n\n", strings.HasSuffix(str, "tring")) // Finding suffix
}
