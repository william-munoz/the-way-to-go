package main

import "fmt"

func main() {
	s, p, d := SumProductDiff(3, 4)
	fmt.Printf("Sum: %d, Product: %d, Difference: %d", s, p, d)
}

func SumProductDiff(i, j int) (s int, p int, d int) { // named version
	s, p, d = i+j, i*j, i-j
	return
}
