package main

import "fmt"

func main() {
	a := 5
	b := 7
	c := 55
	fmt.Println(addNumbers(a, b))
	fmt.Println(subNumbers(a, b))
	fmt.Println(mulNumbers(a, b))
	fmt.Println(c)
}

func addNumbers(a, b int) int {
	return a + b
}

func subNumbers(a, b int) int {
	return a - b
}

func mulNumbers(a, b int) int {
	return a * b
}
