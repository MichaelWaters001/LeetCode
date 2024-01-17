package main

import (
	"fmt"
)

func main() {
	mp := maxProfit([]int{})
	fmt.Printf("%d", mp)
}

type trade struct {
	profit int
	start  int
	end    int
}

//This is a brute force approach TODO optimize 
func maxProfit(prices []int) int {
	l := len(prices)

	//not enough days to make profit
	if l < 2 {
		return 0
	}

	cur := 0
	trades := []trade{}
	max := 0

	//check each possible trade
	for cur < l {
		next := cur + 1
		for next < l {
			t := prices[next] - prices[cur]
			if t > 0 {
				trades = append(trades, trade{profit: t, start: cur, end: next})
				//track max single day trade incase this is better then 2 combined
				if max < t {
					max = t
				}
			}
			next++
		}
		cur++
	}

	//find best 2 non overlapping trades
	for c, t1 := range trades {
		//add all possible combos
		for _, t2 := range trades[c:] {
			//if no overlap
			if t2.start > t1.end {
				p := t1.profit + t2.profit
				if p > max {
					max = p
				}
			}
		}
	}

	return max
}
