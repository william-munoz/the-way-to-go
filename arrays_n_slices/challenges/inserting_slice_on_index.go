package main

import "fmt"

func main() {
	s := []string{"M", "N", "O", "P", "Q", "R"}
	in := []string{"A", "B", "C"}
	res := insertSlice(s, in, 0)
	fmt.Println(res)
	res = insertSlice(s, in, 3)
	fmt.Println(res)
}

func insertSlice(slice, insertion []string, index int) []string {
	ns := make([]string, len(slice)+len(insertion))
	fmt.Println("slice: ", slice[:index])
	at := copy(ns, slice[:index])
	at += copy(ns[at:], insertion)
	copy(ns[at:], slice[index:])
	return ns
}
