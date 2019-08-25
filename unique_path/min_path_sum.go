package unique_path

func minPathSum(grid [][]int) int {
	y := len(grid)
	if y < 1 {
		return 0
	}
	x := len(grid[0])
	if x < 1 {
		return 0
	}

	dp := make([][]int, y)
	for i := 0; i < y; i ++ {
		dp[i] = make([]int, x)
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < y; i ++ {
		for j := 0; j < x; j ++ {
			if i > 0 && j > 0 {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i][j-1] + grid[i][j]
				} else {
					dp[i][j] = dp[i-1][j] + grid[i][j]
				}
			} else if i == 0 && j > 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if i > 0 && j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			}
		}
	}
	return dp[x-1][y-1]
}