package unique_path

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	y := len(obstacleGrid)
	if y < 1 {
		return 1
	}
	x := len(obstacleGrid[0])
	if x < 1 {
		return 1
	}
	fmt.Printf("x=%d y=%d\n", x, y)
	dp := make([][]int, y)
	for i := 0; i < y; i ++ {
		dp[i] = make([]int, x)
	}
	dp[0][0] = 1
	for i := 0; i < y; i ++ {
		for j := 0; j < x; j ++ {
			if obstacleGrid[i][j] > 0 {
				dp[i][j] = 0
			} else if i > 0 && j > 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			} else if i > 0 && j == 0 {
				dp[i][j] = dp[i-1][j]
			} else if i == 0 && j > 0 {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[y-1][x-1]
}