package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := Solution(12)
	fmt.Printf("%d", result)
}

func Solution(A int) int {
	// Implement your solution here
	//handle single digits
	if A < 10 {
		return A
	}

	arr := NumberToArray(A)

	ptr1 := 0
	ptr2 := len(arr) - 1

	result := make([]int, 0)

	for ptr1 < ptr2 {
		result = append(result, arr[ptr1], arr[ptr2])
		ptr1++
		ptr2--
		//handle odd len
		if ptr1 == ptr2 {
			result = append(result, arr[ptr1])
		}
	}
	return ArrayToNumber(result)
}

func NumberToArray(x int) []int {
	arr := make([]int, len(strconv.Itoa(x)))

	for i := len(arr) - 1; x > 0; i-- {
		arr[i] = x % 10
		x = int(x / 10)
	}
	return arr
}

func ArrayToNumber(x []int) int {
	s := ""
	for _, y := range x {
		s = fmt.Sprintf("%s%d", s, y)
	}
	//prod would check err
	i, _ := strconv.Atoi(s)
	return i
}
