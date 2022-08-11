package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "this is an example string"
	fmt.Printf("str value: %s\n", str)
	pointerToStr := strings.NewReader(str)
	fmt.Printf("value that pointerToStr point: %v\n", *pointerToStr)
}
