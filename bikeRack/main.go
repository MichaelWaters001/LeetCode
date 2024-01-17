package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Sprintf("hello world")
	A := []int{10, 0, 8, 2, -1, 12, 11, 3}

	b := Solution(A)

	fmt.Printf("%d\n", b)


	A = []int{10, 0, 8, 2, -12, 12, 11, 3}

	b = Solution(A)

	fmt.Printf("%d\n", b)
}

func Solution(A []int) int {
	// Implement your solution here
	sort.Ints(A)
	largetGap := 0
	l := len(A)
	for i := 1; i < l; i++ {
		curGap := A[i] - A[i-1]
		if curGap > largetGap {
			largetGap = curGap
		}
	}
	return largetGap / 2
}
