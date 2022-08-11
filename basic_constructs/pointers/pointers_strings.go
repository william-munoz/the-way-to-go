package main

import "fmt"

func main() {
	s := "good bye"
	fmt.Printf("Here the original value of s: %s\n", s)
	var p *string = &s
	*p = "ciao" // changing the value at &s

	fmt.Printf("Here is the pointer p: %p\n", p)  // prints memory address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s)   // prints string
}
