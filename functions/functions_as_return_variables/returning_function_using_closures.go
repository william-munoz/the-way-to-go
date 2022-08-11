package main

import "fmt"

func main() {
	// make an Add2 function, give it a name p2, and call it:
	p2 := Add2()
	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))
	// make a special Adder function, a gets value 3:
	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))

	productMessage := ProductMessage(7, "Result of the product of")
	product, message := productMessage(7)
	fmt.Printf("%s is %d\n", message, product)

	var f = NewAdder()
	fmt.Print(f(1), " , ")
	fmt.Print(f(20), " , ")
	fmt.Print(f(300), " , ")
	fmt.Print(f(4000))
}

func Add2() func(b int) int { // return a function
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int { // return a function
	return func(b int) int {
		return a + b
	}
}

// I added this function while I was practicing this topic, that resulted to be very interesting
func ProductMessage(a int, s string) func(b int) (int, string) { // return a function
	return func(b int) (int, string) {
		m := fmt.Sprintf("%s %d times %d", s, a, b)
		return a * b, m
	}
}

func NewAdder() func(int) int {
	var x int // the value the x is retained. Thinking about this, the function is in a variable...
	return func(delta int) int {
		x += delta
		return x
	}
}
