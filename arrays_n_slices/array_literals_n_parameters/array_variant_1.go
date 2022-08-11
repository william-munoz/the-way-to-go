package main

import "fmt"

func main() {
	var arrAge = [5]int{18, 20, 15, 22, 16}
	fmt.Println("Array with fixed size: ", arrAge)
	var arr = [10]int{1, 2, 3}
	fmt.Println("Array with fixed size and some different values and others with default values: ", arr)
}
