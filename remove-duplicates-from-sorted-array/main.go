package main

import "fmt"

func main() {
	a := []int{1, 1, 1, 2, 2, 3}
	removeDuplicates(a)
	fmt.Printf("%+v", a)
}

func removeDuplicates(nums []int) int {
	
	l := len(nums)

	//handle empty 1 or 2 digit array
	if l < 3{
		return l
	}

	count := 2
	cur:=2
	prev2:=0


	for cur<l{
		//if pointers match drop cur
		if nums[prev2] == nums[cur]{
			nums = append(nums[:cur], nums[cur+1:]...)
			l--
			continue
		}
		//update pointers and count
		cur++
		prev2++
		count++
	}

	return count
}
