package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hi, I'm Marc, Hi."
	fmt.Printf("The position of the first instance of \"Marc\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Marc")) // Finding first occurence
	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Hi")) // Finding first occurence
	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi")) // Finding last occurence
	fmt.Printf("The postion of the first instance of \"Burger\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Burger")) // Finding first occurence
}
