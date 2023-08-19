package main

import "fmt"

func main() {
	fmt.Println("starting")

}


//https://leetcode.com/problems/number-of-islands/description/

func numIslands(grid [][]byte) int {
	islandCount := 0
	for y := range grid{
		for x := range grid[y]{
			if grid[y][x] == 1{
				islandCount++
				crawlIsland(grid, y, x)
			}
		}
	}
	return islandCount
}

//mark all adjacent land
func crawlIsland(grid [][]byte, y int, x int){
	
	//return if out of bounds of grid
	if ( y < 0 || y >= len(grid) || x < 0 || x >= len(grid[y])){
		return
	}

	//return if on water
	if grid[y][x] == 0{
		return
	}

	//mark as water so we don't count it again
	grid[y][x]=0

	crawlIsland(grid, y, x-1)
	crawlIsland(grid, y, x+1)
	crawlIsland(grid, y-1, x)
	crawlIsland(grid, y+1, x)
}