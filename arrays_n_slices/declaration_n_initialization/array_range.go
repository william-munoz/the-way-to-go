package main

import "fmt"

func main() {
	var arr1 [5]int // declaring an array
	for i := range arr1 {
		arr1[i] = i * 2 // initializing items of array
		fmt.Printf("Item at index %d has a value %d\n", i, arr1[i])
	}
}
