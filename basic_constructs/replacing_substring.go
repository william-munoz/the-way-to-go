package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "The best programming language of the world is Quechua, and that's why I'm learning Quechua."
	fmt.Printf("%s\n", str)
	fmt.Printf("Sorry. %s\n", strings.Replace(str, "Quechua", "Go", 2))
}
