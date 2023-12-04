package main

import (
	"fmt"
	"math"
)

func main() {
	result := Solution("ABDCA", []int{2}, []int{1})
	fmt.Printf("%d", result)
}

func Solution(S string, X []int, Y []int) int {
	// Implement your solution here

	//map to check for dupes
	points := make(map[rune]float64)

	//track the duplicate closes to the center
	closestDupe := float64(-1)

	for i, tag := range S {
		_, ok := points[tag]
		//dupe tag
		if ok {
			//calc distance
			myDistance := distance(X[i], Y[i])
			//if this is the first dupe tag set min as the further point
			if closestDupe == -1 {
				if myDistance > points[tag] {
					closestDupe = myDistance
				} else {
					closestDupe = points[tag]
					points[tag] = myDistance
				}
			} else if myDistance < closestDupe { //keep further of old closest or point being overwritten
				closestDupe = points[tag]
				points[tag] = myDistance

			}
			//drop any further points

		} else { // new tag
			points[tag] = distance(X[i], Y[i])
		}
	}

	//if there are no dup tags all count
	if closestDupe < 0 {
		return len(points)
	}

	count := 0
	//count all points closer then closest dupe
	for _, point := range points {
		if point < closestDupe {
			count++
		}
	}

	return count
}

// distance from 0,0
func distance(a int, b int) float64 {
	return (math.Sqrt(math.Pow(float64(a), 2) + math.Pow(float64(b), 2)))
}
