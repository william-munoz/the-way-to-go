package main

import "fmt"

func main() {
	// if s is a string, so also an array of bytes
	s := "nirvana"
	c := []byte(s)
	fmt.Printf("value of s: %s\n", s)
	fmt.Printf("value of c: %v\n", c)

	// this can also be done whit the copy function
	dst := []byte("")
	src := "professional go developer"
	copy(dst, src)
	fmt.Printf("value of dst: %v\n", dst)
	fmt.Printf("value of src: %v\n", src)

	// the for range can also be applied as in the following program
	s = "\u00ff\u754c"
	for i, c := range s {
		fmt.Printf("new values of s: %d:%c\n", i, c)
	}

	// a string may also be appended to a byte slice
	var b []byte
	s = "professional go developer"
	b = append(b, s...)
	fmt.Printf("new value of s: %s\n", s)
	fmt.Printf("value of b: %v\n", b)

	// show how to make substrings from a string
	showMakingSubstring()
	showChagingCharacterString()
}

func showMakingSubstring() {
	s := "professional go developer"
	substring := s[:]
	fmt.Printf("value of s: %s\n", s)
	fmt.Printf("substring 1: %s\n", substring)
	fmt.Printf("substring 2: %s\n", s[:len(s)-10])
	fmt.Printf("substring 3: %s\n", s[13:])
	fmt.Printf("substring 4: %s\n", s[13:15])
}

func showChagingCharacterString() {
	// strings are immutable
	// to change a character, you first have to convert the string to an array of bytes.
	// then, an array-item of certain index can be changed, and the array must be converted bact to a new string.
	// for example, hello to hallo
	s := "hello"
	c := []byte(s)
	c[1] = 'a'
	s2 := string(c)
	fmt.Printf("new string s2: %s\n", s2)
}
