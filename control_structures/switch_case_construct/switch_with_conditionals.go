package main

import "fmt"

func main() {
	num1 := 100
	switch {
	case num1 < 0:
		fmt.Println("Number if negative")
	case num1 > 0 && num1 < 10:
		fmt.Println("Number is between 0 and 10")
	default:
		fmt.Println("Number is 10 or greater")
	}
}
