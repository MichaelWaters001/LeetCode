package main

import "fmt"

func main() {
	inputHeight := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(inputHeight))
}

// to slow, but works. only taller walls matter when starting with max width
// func maxArea(height []int) int {
//     maxArea := 0
// 	for x1, y1 := range height {
// 		for x2, y2 := range height[x1+1:] {
// 			area := (x2+1) * Min(y1, y2)
// 			if area > maxArea {
// 				maxArea = area
// 			}
// 		}
// 	}
// 	return maxArea
// }

// func Min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }


func maxArea(height []int) int {
	size := len(height)
	left := 0
	right := size - 1
	maxWidth := size - 1
	maxArea := 0

	for width := maxWidth; width > 0; width-- {
		if height[left] < height[right] {
			area := width * height[left]
			if area > maxArea {
				maxArea = area
			}
			left++
		} else {
			area := width * height[right]
			if area > maxArea {
				maxArea = area
			}
			right--
		}
	}
	return maxArea
}