package main

import "fmt"

func main() {
	fmt.Println("In main function before calling greeting function.")

	greeting()

	fmt.Println("In main function after calling greeting function.")
}

func greeting() {
	fmt.Println("In greeting function: Hi!!!")
}
