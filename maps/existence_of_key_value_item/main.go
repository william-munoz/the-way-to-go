package main

import "fmt"


var programmingLanguages = map[string]string{
	"golang": "Go Programming Language",
	"python": "Python Programming Language",
	"java": "Java Programming Language",
	"swift": "Swift Programming Language",
	"rust": "Rust Programming Language",
	"c": "C Programming Language",
}

func main() {
	originalContentOfProgrammingLanguage()
	testingExistenceOfAKey()
	deletingAnElementWithAKey()
	originalContentOfProgrammingLanguage()
	theForRangeConstruct()
}

func originalContentOfProgrammingLanguage() {
	fmt.Println("")
	fmt.Println("--- Content of programmingLanguages variable ---")
	fmt.Printf("programmingLanguages values: %v\n", programmingLanguages)
	fmt.Println("")
}

func testingExistenceOfAKey() {
	fmt.Println("")
	fmt.Println("--- Testing existence of a key ---")
	pl, ok := programmingLanguages["golang"] 
	if ok {
		fmt.Printf("the variable programmingLanguages containt a key \"golang\" and its value is: %s", pl)
	}
}

func deletingAnElementWithAKey() {
	fmt.Println("")
	fmt.Println("--- Deleting an element with a key ---")
	delete(programmingLanguages, "java")
	fmt.Println("deleting the ugly java programming language")
}

func theForRangeConstruct() {
	fmt.Println("")
	fmt.Println("--- The for range construct ---")
	for key, value := range programmingLanguages {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}
}
