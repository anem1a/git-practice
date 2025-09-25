package main

import "fmt"

func main() {
	a := 5
	b := 7
	c := 3
	fmt.Println(addNumbers(a, b))
	fmt.Println(c)
}

func addNumbers(a, b int) int {
	return a + b
}