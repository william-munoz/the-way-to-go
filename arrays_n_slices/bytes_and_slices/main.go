package main

import (
	"bytes"
	"fmt"
)

func main() {
	// new buffer
	var b bytes.Buffer

	// write strings to the buffer
	b.WriteString("ABC")
	b.WriteString("DEF")
	fmt.Println(b.String())

	// use Fprintf with buffer
	fmt.Fprintf(&b, " A number: %d, a string: %v\n", 10, "bird")
	fmt.Println(b.String())

	// converting to a string and print it
	b.WriteString("[DONE]")
	fmt.Println(b.String())

	// writing more values to a buffer
	b.WriteString("something else")
	fmt.Println(b.String())

	// identify the final value of b
	fmt.Println()
	fmt.Println("Final value of b:")
	fmt.Println(b.String())
}
