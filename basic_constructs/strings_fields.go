package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "this is an example string"
	fmt.Printf("Original str value: %s\n", str)
	strList := strings.Fields(str)
	fmt.Printf("Resulted fields: %v\n", strList)
	fmt.Printf("strList count: %d\n", len(strList))
	fmt.Printf("First value of strList: %s\n", strList[0])
	fmt.Printf("Last value of strList: %s\n", strList[len(strList)-1])
}
