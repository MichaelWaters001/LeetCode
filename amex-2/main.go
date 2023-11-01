package main

import "sort"

func main() {
	print("find unused least positive int")
	Solution([]int{1, 3, 6, 4, 1, 2})
}

func Solution(A []int) int {
	//not enough sides given
	if len(A) < 3 {
		return -1
	}
	per := -1
	sort.Ints(A)

	for i := 0; i < len(A)-2; i++ {
		if validTri(A[i],A[i+1],A[i+2]){
			per = A[i]+A[i+1]+A[i+2]
		}
	}
	return per

}

func validTri(p int, q int, r int) bool {
	if p+q <= r {
		return false
	}
	if q+r <= p {
		return false
	}
	if r+p <= q {
		return false
	}
	return true
}