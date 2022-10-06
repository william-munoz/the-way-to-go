package main

import "fmt"

func main() {
	fmt.Println("--- Maps ---")

	// the length of the map doesn't have to be known at the declaration, which means a map can grow dynamically.
	// the value of an uninitialized map is nil.
	// the key type can be any type for which the operations == and != are defined.
	// looking up a value in map by key is fast, much faster than a linear search.
	// but still around 100x slower than direct indexing in an array or slice.
	// if performance is very important try to solve the problem with slices.
	
	mapLit := map[string]int{"one": 1, "two": 2}
	var mapAssigned map[string]int
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit
	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3

	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapAssigned["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])
	fmt.Println("")

	fmt.Println("--- Map storing a function as value type ---")
	storeFunctionsInMap()

	mapCapacity()
	slicesAsMapValues()
}

func storeFunctionsInMap() {
	mf := map[int]func() int{ // key type int, and value type func()int
		1: func() int { return 10 },
		2: func() int { return 20 },
		3: func() int { return 30 },
	}
	fmt.Println(mf)
}

func mapCapacity() {
	mf := make(map[string]float32, 100)
	fmt.Println("")
	fmt.Println("--- Map Capacity ---")
	fmt.Printf("map initial value: %v\n", mf)
	fmt.Printf("map length: %d", len(mf))

	noteFrequence := map[string]float32{
		"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83, "G0": 24.50, "A0": 27.50, "B0": 30.87, "A4": 440,
	}
	fmt.Println("")
	fmt.Println("--- Note Frequencies ---")
	fmt.Printf("note frequencies: %v\n", noteFrequence)
}

func slicesAsMapValues() {
	fmt.Println("")
	fmt.Println("--- Slices as map values ---")
	ms1 := make(map[int][]int)
	ms2 := make(map[int]*[]int)
	fmt.Printf("Initial value of ms1: %v\n", ms1)
	fmt.Printf("Initial value of ms2: %v\n", ms2)
}
