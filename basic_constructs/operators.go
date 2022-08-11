package main

import "fmt"

func main() {
	arithmeticOperators()
	logicalOperators()
}

func arithmeticOperators() {
	b3 := 10 > 5 // greater than operator
	fmt.Println(b3)
	b3 = 10 < 5 // less than operator
	fmt.Println(b3)
	b3 = 5 <= 10 // less than equal to operator
	fmt.Println(b3)
	b3 = 10 != 10 // not equal to operator
	fmt.Println(b3)
}

func logicalOperators() {
	b3 := 10 > 5 && 7 < 15 // AND operator
	fmt.Println(b3)
	b3 = 10 < 5 || 1 > 7 // OR operator
	fmt.Println(b3)
	b3 = !b3 // NOT operator
	fmt.Println(b3)
}
