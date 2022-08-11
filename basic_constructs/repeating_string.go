package main

import (
	"fmt"
	"strings"
)

func main() {
	origs := "Hi there! "
	news := strings.Repeat(origs, 3)
	fmt.Printf("The new repeated string is: %s\n", news)
}
