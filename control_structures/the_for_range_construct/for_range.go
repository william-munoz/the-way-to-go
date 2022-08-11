package main

import "fmt"

func main() {
	str := "Go is a beautiful language!"

	// for range
	for pos, char := range str {
		fmt.Printf("Character on position %d is: %c\n", pos, char)
	}
	fmt.Println()
	str2 := "Chinese: 日本語"

	// for range
	for pos, char := range str2 {
		fmt.Printf("Character %c starts at byte position %d\n", char, pos)
	}

	fmt.Printf("\nWe can see that the normal English characters are represented by 1 byte, and the Chinese characters by 3 bytes.\n")
}
