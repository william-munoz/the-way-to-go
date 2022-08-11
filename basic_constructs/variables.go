package main

import "fmt"

func main() {
	variableDefaultValue()
	variableExplicitValue()
	variableAutomaticTypeInference()
	variableShortOnlyFunctions()
	variableAssignAndDeclare()
}

func variableDefaultValue() {
	var number int
	fmt.Println(number)
	var decision bool
	fmt.Println(decision)
}

func variableExplicitValue() {
	var number int = 7
	fmt.Println(number)
	var decision bool = true
	fmt.Println(decision)
}

func variableAutomaticTypeInference() {
	var number = 7
	fmt.Println(number)
	var decision = true
	fmt.Println(decision)
}

func variableShortOnlyFunctions() {
	number := 7
	fmt.Println(number)
	decision := true
	fmt.Println(decision)
}

func variableAssignAndDeclare() {
	number := 7
	fmt.Println(number)
	number, decision := 77, true
	fmt.Println(number, decision)
}
