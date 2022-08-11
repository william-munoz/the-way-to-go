package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	fmt.Printf("a array is: %d\n", a)
	a = enlarge(a, 5)
	fmt.Printf("new a array is: %d\n", a)
	fmt.Printf("new capacity of a: %d\n", cap(a))
}

func enlarge(s []int, factor int) []int {
	ns := make([]int, len(s)*factor)
	copy(ns, s)
	return ns
}
