package main

import (
	"fmt"
)

func main() {
	result := Solution([]int{1, 0, 2, 4}, []int{2, 2, 0, 5}, []int{3, 17, 7, 4, 5, 17})
	fmt.Printf("%d", result)
}

func Solution(start []int, dest []int, limit []int) int {
	// Implement your solution here

	//total cost for all trips
	cost := 0
	for i := 0; i < len(start); i++ {
		cost = cost + stopCost(start[i], dest[i])
	}

	max := maxCost(start, dest, limit)
	if cost < max{
		return cost
	}
	return max
}

//cost for single trip
func stopCost(a int, b int) int {
	val := a - b
	if val < 0 {
		return (val*-1)*2 + 1
	}
	return val*2 + 1
}

func maxCost(a []int, b []int, max []int) int {
	maxStation := 0
	//get the highest station 
	for _, station := range a {
		if station > maxStation {
			maxStation = station
		}
	}
	//dest may be the highest
	for _, station := range b {
		if station > maxStation {
			maxStation = station
		}
	}
	//return max value of highest station visited
	return max[maxStation]
}
