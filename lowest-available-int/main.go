package main

import "sort"

func main() {
	print("find unused least positive int")
	Solution([]int{1, 3, 6, 4, 1, 2})
}

func Solution(A []int) int {
	sort.Ints(A)
	low := 1
	for _, num := range A {
		//skip past negative values
		if num < 1 {
			continue
		}

		if low < num {
			return low
		}
		low = num + 1
	}
	return low
}
