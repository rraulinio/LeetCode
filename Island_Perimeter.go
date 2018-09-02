func islandPerimeter(grid [][]int) int {
	var total int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				tmp := 4
				if i-1 >= 0 && grid[i-1][j] == 1 { // up
					tmp--
				}
				if i+1 < len(grid) && grid[i+1][j] == 1 { // down
					tmp--
				}
				if j-1 >= 0 && grid[i][j-1] == 1 { // left
					tmp--
				}
				if j+1 < len(grid[0]) && grid[i][j+1] == 1 { // right
					tmp--
				}
				total += tmp
			}
		}
	}
	return total
}