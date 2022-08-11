package main

import "fmt"

func main() {
	var i1 int
	var f1 float32
	i1, _, f1 = ThreeValues()
	fmt.Printf("The int: %d, the float: %.2f\n", i1, f1)
}

func ThreeValues() (int, int, float32) {
	return 5, 6, 7.5
}
