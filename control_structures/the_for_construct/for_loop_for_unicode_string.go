package main

import "fmt"

func main() {
	str := "Go is beautiful language!"
	fmt.Printf("The lenght of str is: %d\n", len(str))

	// for loop to find character at each position
	for i := 0; i < len(str); i++ {
		fmt.Printf("Character on position %d is: %c \n", i, str[i])
	}
}
