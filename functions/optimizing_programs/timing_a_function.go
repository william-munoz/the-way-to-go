package main

import (
	"fmt"
	"time"
)

func Calculation() {
	for i := 0; i < 10000; i++ {
		//do something
	}
}
func main() {
	start := time.Now()
	Calculation()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("Calculation took this amount of time: %s\n", delta)
}
