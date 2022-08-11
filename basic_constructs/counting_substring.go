package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "The best programming language of the world is Go, and that's why I'm learning Go."
	fmt.Printf("%s\n", str)
	fmt.Printf("Number of times that Go appears on str variable: %d\n", strings.Count(str, "Go"))
}
