package main

import "fmt"

func main() {
	characterType()
	characterUnicode()
}

func characterType() {
	var ch1 byte = 'A'
	var ch2 byte = 65
	var ch3 byte = '\x41'
	fmt.Printf("ch1 value is: %v\n", ch1)
	fmt.Printf("ch2 value is: %v\n", ch2)
	fmt.Printf("ch3 value is: %v\n", ch3)
}

func characterUnicode() {
	var ch1 int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch1, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch1, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch1, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U\n", ch1, ch2, ch3) // UTF-8 code point
}
