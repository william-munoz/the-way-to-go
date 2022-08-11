package main

import "fmt"

var number int = 7
var score int

func main() {
	variableGlobalScope()
	variableLocalScope()
}

func variableGlobalScope() {
	fmt.Println("Original value of global number variable: ", number)
	number = 77
	fmt.Println("New value of global number variable: ", number)
	fmt.Println("Initial value of global score variable: ", score)
	score = 7
	fmt.Println("New value of global score variable: ", score)
}

func variableLocalScope() {
	var decision = true
	fmt.Println("Value of local variable decision: ", decision)
}
