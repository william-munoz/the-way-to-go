package main

import (
	"container/list"
	"fmt"
)

func main() {
	lst := insertListElements(5)

	fmt.Printf("list returned: %v\n", lst)

	for e := lst.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func insertListElements(n int) *list.List {
	lst := list.New()

	for i := 1; i <= n; i++ {
		lst.PushBack(i)
	}

	return lst
}
