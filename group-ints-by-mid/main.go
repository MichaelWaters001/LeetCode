package main

import "fmt"

func main() {
	list := []int{18123451111, 90123452222, 28123451234, 10123461111, 46234561111, 65334561111, 90987654321}

	solution(list)
}

func solution(list []int) (groups map[int]int) {
	groups = make(map[int]int)

	for _, num := range list{
		num = num % 1000000000
		num = num / 10000
		groups[num]++
	} 

	fmt.Printf("%+v",groups)

	return
}
