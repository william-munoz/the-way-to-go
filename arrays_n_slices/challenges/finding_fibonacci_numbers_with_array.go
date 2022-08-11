package main

import "fmt"

var fib [10]int64

func main() {
	fmt.Printf("Fibonacci: %v\n", fibs())
}

func fibs() [10]int64 {
	for i := 0; i < len(fib); i++ {
		if i <= 1 {
			fib[i] = 1
		} else {
			fib[i] = fib[i-2] + fib[i-1]
		}
	}

	return fib
}
