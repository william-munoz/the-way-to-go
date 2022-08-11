package main

import "fmt"

func main() {
	b()
}

func trace(s string) {
	fmt.Println("Entering: ", s)
}

func untrace(s string) {
	fmt.Println("Leaving: ", s)
}

func a() {
	trace("a")
	defer untrace("a") // untracing via defer
	fmt.Println("in a")
}

func b() {
	trace("b")
	defer untrace("b") // untracing via defer
	fmt.Println("in b")
	a()
}
