package main

import (
	"fmt"
	"strconv"
)


func main() {
	a := Solution(1)
	fmt.Println(a)

	a = Solution(11)
	fmt.Println(a)

	a = Solution(200)
	fmt.Println(a)
}

func Solution(N int) int {
	// Implement your solution here
	s := strconv.Itoa(N)
	l := len(s)
	if l == 1 {
		return 0
	}
	nString := "1"
	for i := 1; i < l; i++ {
		nString = nString + "0"
	}
	n, _ := strconv.Atoi(nString)
	return n
}
