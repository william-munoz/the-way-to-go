package main

import "fmt"

var Days = map[int]string{
	1: "Monday",
	2: "Thuesday",
	3: "Wednesday",
	4: "Thursday",
	5: "Friday",
	6: "Saturday",
	7: "Friday",
}

func findDay(n int) string {
	val, isPresent := Days[n]
	if isPresent {
		return val
	}
	return "Wrong key"
}

func main() {
	n := 4
	fmt.Println(n, ":", findDay(n))
	n = 0
	fmt.Println(n, ":", findDay(n))
}
