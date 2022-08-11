package main

import "fmt"

func main() {
	var arrLazy = [...]int{5, 6, 7, 8, 22}
	// ... indicates the compiler has to count the number of items to obtain the lenght of the array.
	fmt.Printf("Array with fixed size of lenght %d and values %v\n", len(arrLazy), arrLazy)
	// if the ... is omitted then a slice is created.

	// [...]int is NOT a type.
}
