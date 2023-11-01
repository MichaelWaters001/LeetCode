package main

import "fmt"

func main() {
	result := add(1, 2)
	fmt.Printf("%d", result)
}

func add(a int, b int) int {
	return a + b
}
