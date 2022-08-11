package main

import (
	"fmt"
	"strings"
)

const sep = "-"

func main() {
	str := "this-is-an-example-string"
	fmt.Printf("Original str: %s\n", str)
	strList := strings.Split(str, sep)
	fmt.Printf("strList value: %v\n", strList)
	fmt.Printf("First strList's item: %s\n", strList[0])
	fmt.Printf("Lst strList's item: %s\n", strList[len(strList)-1])
}
