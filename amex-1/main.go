package main

import "fmt"

func main() {
	print("find unused least positive int")
	bad := find_min(Solution(1000))
	fmt.Print(bad)
}

func find_min(A []int) int {
	var result int = 0
	for i := 1; i < len(A); i++ {
		if result > A[i] {
			result = A[i]
		}
	}
	return result
}

func Solution(A int) []int {
	var ary []int
	for i := 1; i <= A; i++{
		ary = append(ary,i)
	}
	return ary
}
