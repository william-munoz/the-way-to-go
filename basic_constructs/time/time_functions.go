package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())

	timeFormat()
}

func timeFormat() {
	t := time.Now().UTC()
	fmt.Println(t.Format("02 Jan 2006 15:04"))
}
