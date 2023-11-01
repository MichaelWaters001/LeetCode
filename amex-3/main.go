package main

import "sort"

func main() {
	Solution([]int{1, 2, 3, 5})
}

func Solution(A []int) int {
	// Implement your solution here
	sort.Ints(A)
	possibility := 1
	//handle empty
	if len(A) == 0 {
		return 1
	}
	for _, cur := range A {
		//skip over negative numbers
		if cur < 1 {
			continue
		}
		//if available
		if possibility < cur {
			return possibility
		}
		// set = do not inc to account for repeated numbers
		possibility = cur + 1
	}
	//no gaps until end of array
	return possibility
}
