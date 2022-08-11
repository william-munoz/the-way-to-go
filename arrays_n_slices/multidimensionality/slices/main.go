package main

import "fmt"

func main() {
	values := [][]int{} // multidimensional slice

	// these are the first two rows
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	row3 := []int{7, 8, 9, 10}

	// append each row to the two-dimensional slice
	values = append(values, row1)
	values = append(values, row2)
	values = append(values, row3) // rows could be different sizes

	// display first row
	fmt.Println("Row 1")
	fmt.Println(values[0])

	// display second row
	fmt.Println("Row 2")
	fmt.Println(values[1])

	// access an element
	fmt.Println("First element from first row")
	fmt.Println(values[0][0])

	// first element from second row
	fmt.Println("First element from second row")
	fmt.Println(values[1][0])

	// last element from third row
	fmt.Println("Last element from third row")
	fmt.Println(values[2][len(row3)-1])

	// display entire slice
	fmt.Println("Values")
	fmt.Println(values)
}
