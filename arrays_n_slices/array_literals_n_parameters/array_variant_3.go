package main

import "fmt"

func main() {
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	fmt.Printf("Array with key-value: %v\n", arrKeyValue)
	fmt.Printf("Value of first element: %s\n", arrKeyValue[0])
}
