package main

import "fmt"

func main() {
	n := 0
	reply := &n
	fmt.Println("Before multiplication: ", n)
	fmt.Println("Before multiplication: ", *reply)

	Multiply(10, 5, reply)
	fmt.Println("Multiply: ", *reply) // Multiply: 50
	fmt.Println("Multiply: ", n)      // Multiply: 50
}

// this function changes reply:
func Multiply(a, b int, reply *int) {
	*reply = a * b // side-effect(changing n)
}
