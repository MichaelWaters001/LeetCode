package main

//https://leetcode.com/problems/check-if-there-is-a-valid-partition-for-the-array/

func validPartition(nums []int) (valid bool) {

	return test(nums, 0)
}

// ptr tracks how far we can traverse the slice
func test(cur []int, ptr int) bool {

	//reached the end evenly success
	if ptr == len(cur) {
		return true
	}

	if (ptr + 3) <= len(cur) { //overflow test
		if threeOrdered(cur[ptr : ptr+3]) {
			if test(cur, ptr+3) {
				return true
			}
		}
		if threeSame(cur[ptr : ptr+3]) {
			if test(cur, ptr+3) {
				return true
			}
		}
	}
	if (ptr + 2) <= len(cur) { //overflow test
		if twoSame(cur[ptr : ptr+2]) {
			if test(cur, ptr+2) {
				return true
			}
		}
	}
	return false
}

func twoSame(a []int) bool {
	return a[0] == a[1]
}

func threeSame(a []int) bool {
	return a[0] == a[1] && a[1] == a[2]
}

func threeOrdered(a []int) bool {
	return a[0] == (a[1]-1) && a[0] == (a[2]-2)
}
