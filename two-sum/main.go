package main

//https://leetcode.com/problems/two-sum/description/

func twoSum(nums []int, target int) []int {

	// m [num]og index
	m := make(map[int]int)
	for i, num := range nums {
		//calculate the needed second value
		need := target - num
		//test if we have seen it
		if val, ok := m[need]; ok {
			return []int{val, i}
		}

		//store current value
		m[num] = i

	}
	return nil
}
