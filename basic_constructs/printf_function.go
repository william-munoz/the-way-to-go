package main

import "fmt"

var number = 7

func main() {
	printGlobalVariable()
	printBoolVariable()
}

func printGlobalVariable() {
	fmt.Printf("Original value of global number variable: %d\n", number)
}

func printBoolVariable() {
	var decision = true
	fmt.Printf("Value of decision variable: %t\n", decision)
}
