package main

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
	t = time.Now().UTC()
	fmt.Println(t)
	fmt.Println(time.Now())
	// calculating times:
	week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
	weekFromNow := t.Add(week)
	fmt.Println(weekFromNow)
	// formatting times:
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("02 Jan 2006 15:04"))
	s := t.Format("2006 01 02")
	fmt.Println(t, "=>", s)
}
