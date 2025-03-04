package main

import "fmt"

func main() {
	gain := []int{-5, 1, 5, 0, -7}
	fmt.Println(largestAltitude(gain))
}

func largestAltitude(gain []int) int {
	max := 0
	cur := 0
	for _, change := range gain{
		cur += change
		if cur > max{
			max = cur
		}
	}

	return max
}