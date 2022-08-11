package main

import "fmt"

func main() {
	var c1 complex64 = 5 + 10i
	fmt.Printf("The value of c1 is: %v\n", c1)
	fmt.Printf("Real part of c1 is: %v\n", real(c1))
	fmt.Printf("Imaginary part of c1 is: %v\n", imag(c1))
}
