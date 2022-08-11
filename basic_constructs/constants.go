package main

import (
	"fmt"
)

func main() {
	implicit()
	explicit()
	typed()
	compilation()
	multipleAssignments()
	enumerations()
	enumerationsWithIOTA()
}

func implicit() {
	// Remark: There is a convention to name constant identifiers will all uppercase letters
	const PI = 3.14159
	fmt.Println(PI)
}

func explicit() {
	const B string = "hello"
	fmt.Println(B)
}

func typed() {
	var n int
	x := n + 5 // untyped numeric constant "5" becomes typed as int, because n was int.
	fmt.Println(x)

	const MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY int = 1, 2, 3, 4, 5, 6 // typed constants because the type is mentioned
}

func compilation() {
	const C1 = 2 / 3 // value can be inferred at compilation time
	// const C2 = getNumber() // value cannot be inferred at compile-time.
	// Constant's value should be known at compile time according to the design principles
}

func multipleAssignments() {
	const BEEF, TWO, C = "meat", 2, "veg" // all constants are untyped.
}

func enumerations() {
	const (
		UNKNOWN = 0
		FEMALE  = 1
		MALE    = 2
	)
}

func enumerationsWithIOTA() {
	type Gender int
	const (
		UNKNOWN = iota
		FEMALE
		MALE
	)
}
