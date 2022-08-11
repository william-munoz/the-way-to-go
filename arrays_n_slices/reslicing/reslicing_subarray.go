package main

import "fmt"

func main() {
	var ar = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var a = ar[5:7]
	fmt.Println(a)
	fmt.Printf("a array lenght: %d\n", len(a))
	fmt.Printf("a array capacity: %d\n", cap(a))

	a = a[0:4] // ref of subarray {5,6,7,8}
	fmt.Println(a)
	fmt.Printf("a array lenght: %d\n", len(a))
	fmt.Printf("a array capacity: %d\n", cap(a))

	a = a[0:5] // ref of subarray {5,6,7,8,9}
	fmt.Println(a)
	fmt.Printf("a array lenght: %d\n", len(a))
	fmt.Printf("a array capacity: %d\n", cap(a))

	a = append(a, 10)
	fmt.Println(a)
	fmt.Printf("a array lenght:%d\n", len(a))
	fmt.Printf("a array capacity: %d\n", cap(a))

	a = a[0:6]
	a[5] = 11
	fmt.Println(a)
	fmt.Printf("a array lenght:%d\n", len(a))
	fmt.Printf("a array capacity: %d\n", cap(a))
}
