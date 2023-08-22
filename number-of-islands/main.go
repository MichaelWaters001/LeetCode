package main

import "fmt"

func main() {
	fmt.Println("starting")

}


//https://leetcode.com/problems/number-of-islands/description/

func numIslands(grid [][]byte) int {
	islandCount := 0
	for x, row := range grid{
		for y, cell := range row{
			if cell == '1'{
				islandCount++
				crawlIsland(grid, x, y)
			}
		}
	}
	return islandCount
}

//mark all adjacent land
func crawlIsland(grid [][]byte, x int, y int){
	
	//return if out of bounds of grid
	if ( x < 0 || x >= len(grid) || y < 0 || y >= len(grid[x]) ){
		return
	}

	//return if on water
	if grid[x][y] == '0'{
		return
	}

	//mark as water so we don't count it again
	grid[x][y]='0'

	crawlIsland(grid, x-1, y)
	crawlIsland(grid, x+1, y)
	crawlIsland(grid, x, y-1)
	crawlIsland(grid, x, y+1)
}