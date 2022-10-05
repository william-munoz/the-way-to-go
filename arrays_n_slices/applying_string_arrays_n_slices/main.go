package main

import (
	"fmt"
	"sort"
)

func main() {
	// if s is a string, so also an array of bytes
	fmt.Println("--- Array of bytes ---")
	s := "nirvana"
	c := []byte(s)
	fmt.Printf("value of s: %s\n", s)
	fmt.Printf("value of c: %v\n", c)
	fmt.Println("")

	// this can also be done whit the copy function
	fmt.Println("--- Copy funtion ---")
	dst := []byte("")
	src := "professional go developer"
	copy(dst, src)
	fmt.Printf("value of dst: %v\n", dst)
	fmt.Printf("value of src: %v\n", src)
	fmt.Println("")

	// the for range can also be applied as in the following program
	fmt.Println("--- Range in a string ---")
	s = "\u00ff\u754c"
	for i, c := range s {
		fmt.Printf("new values of s: %d:%c\n", i, c)
	}
	fmt.Println("")

	// a string may also be appended to a byte slice
	fmt.Println("--- Append values to string ---")
	var b []byte
	s = "professional go developer"
	b = append(b, s...)
	fmt.Printf("Current value of s: %s\n", s)
	fmt.Printf("value of b: %v\n", b)
	fmt.Println("")

	// show how to make substrings from a string
	showMakingSubstring()

	// changing a character in a string
	showChangingCharacterString()

	// searching and sorting slices and arrays
	searchAndSortingSlicesNArrays()

	// non-consecutive
	fmt.Println("--- non-consecutive values in array ---")
	arr := []byte{'a', 'b', 'a', 'a', 'a', 'c', 'd', 'e', 'f', 'g'}
	fmt.Printf("value of arr: %v\n", arr)
	for _, c := range arr {
		fmt.Printf("%c ", c)
	}
	nonConArr := nonConsecutive(arr)
	fmt.Printf("\nvalue of non-consecutive: %v\n", nonConArr)
	for _, c := range nonConArr {
		fmt.Printf("%c ", c)
	}
	fmt.Println("")

}

func showMakingSubstring() {
	s := "professional go developer"
	substring := s[:]
	fmt.Println("--- Making a substring of string ---")
	fmt.Printf("value of s: %s\n", s)
	fmt.Printf("substring 1: %s\n", substring)
	fmt.Printf("substring 2: %s\n", s[:len(s)-10])
	fmt.Printf("substring 3: %s\n", s[13:])
	fmt.Printf("substring 4: %s\n", s[13:15])
	fmt.Println("")
}

func showChangingCharacterString() {
	// strings are immutable
	// to change a character, you first have to convert the string to an array of bytes.
	// then, an array-item of certain index can be changed, and the array must be converted bact to a new string.
	// for example, hello to hallo
	s := "hello"
	c := []byte(s)
	c[1] = 'a'
	s2 := string(c)
	fmt.Println("Change character string")
	fmt.Printf("new string s2: %s\n", s2)
	fmt.Println("")
}

func searchAndSortingSlicesNArrays() {
	a := []int{107, 7, 97, 77, 27, 37, 47, 67, 17, 87, 57}
	fmt.Println("--- Searching and sorting slices and arrays ---")
	fmt.Printf("value of a: %v\n", a)
	if sort.IntsAreSorted(a) == false {
		fmt.Println("the a array is not sorted")
	}
	sort.Ints(a)
	fmt.Printf("value of a sorted: %v\n", a)
	if sort.IntsAreSorted(a) {
		fmt.Println("the a array is sorted")
	}

	s := []string{"Alejo", "Tali", "Magally", "Will", "Leo", "Abuela", "Abuelo", "Tato"}
	fmt.Printf("value of s: %v\n", s)
	sort.Strings(s)
	fmt.Printf("value of s sorted: %v\n", s)
	fmt.Println("")

	// searching in an array it must be sorted first
	i := sort.SearchInts(a, 77)
	fmt.Printf("position where the number %v is located: %d\n", a[i], i)
	fmt.Println("")

	i = sort.SearchStrings(s, "Will")
	fmt.Printf("position where the string %v is located: %d\n", s[i], i)
	fmt.Println("")
}

func nonConsecutive(s []byte) []byte {
	arru := make([]byte,len(s))
	ixu := 0
	tmp := byte(0)
	for _, val := range s {
		if val != tmp {
			arru[ixu] = val
			ixu++
		}
		tmp = val
	}
	return arru
}
