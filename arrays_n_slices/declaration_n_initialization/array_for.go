package main

import "fmt"

func main() {
	var arr1 [5]int // declaring the array
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2 // initializing items of array
	}
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("Item at index %d has a value %d\n", i, arr1[i]) // printing items of array
	}
}
