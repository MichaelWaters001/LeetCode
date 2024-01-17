package main

import "fmt"

func main() {
	mp := maxProfit([]int{})
	fmt.Printf("%d", mp)
}

func maxProfit(prices []int) int {
	profit := 0

	l := len(prices)

	//not enough days to make profit
	if l < 2 {
		return 0
	}

	cur := 1
	//low := prices[0]

	//check each day
	for cur < l {
		//buy if yesterday was lower, does not matter if we by and sell same day
		if prices[cur-1] < prices[cur] {
			profit += (prices[cur] - prices[cur-1])
		}
		cur++
	}

	return profit
}
