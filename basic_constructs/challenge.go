package main

import (
	"fmt"
)

type Celsius float32
type Fahrenheit float32

func main() {
	f := toFahrenheit(100)
	fmt.Printf("100 in Fahrenheit is: %v\n", f)
}

func toFahrenheit(t Celsius) Fahrenheit {
	var temp Fahrenheit
	temp = Fahrenheit((t * 9 / 5) + 32)
	return temp
}
