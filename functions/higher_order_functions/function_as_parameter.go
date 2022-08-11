package main

import "fmt"

func main() {
	callback(1, Add)
	callback(7, Product)
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func Product(a, b int) {
	fmt.Printf("The product of %d and %d is: %d\n", a, b, a*b)
}

func callback(y int, f func(int, int)) {
	f(y, 2) // this becomes Add(y, 2) or Product(y, 2)
}
