package algorithms

func NumIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				explore_island(grid, i, j)
				count++
			}
		}
	}
	return count
}

func explore_island(grid [][]byte, i, j int) {
	if i < len(grid) && i >= 0 && j < len(grid[0]) && j >= 0 {
		if grid[i][j] == '1' {
			grid[i][j] = '0'
			explore_island(grid, i, j+1)
			explore_island(grid, i+1, j)
			explore_island(grid, i-1, j)
			explore_island(grid, i, j-1)
		}
	}
}
