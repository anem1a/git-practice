package main

import "fmt"

func main() {
	a := 5
	b := 7
	fmt.Println(addNumbers(a, b))
}

func addNumbers(a, b int) int {
	return a + b
}