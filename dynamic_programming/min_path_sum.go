// Given a m n n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.

// Note: mou can only move either down or right at any point in time.

// Example:

// Input:
// [
//   [1,3,1],
//   [1,5,1],
//   [4,2,1]
// ]
// Output: 7
// Explanation: Because the path 1→3→1→1→1 minimizes the sum.

package dynamic_programming

import (
// "math"
)

func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				if j > 0 {
					grid[i][j] += grid[i][j-1]
				}
			} else {
				if j == 0 {
					grid[i][j] += grid[i-1][j]
				} else if grid[i-1][j] > grid[i][j-1] {
					grid[i][j] += grid[i][j-1]
				} else {
					grid[i][j] += grid[i-1][j]
				}
			}
		}
	}

	return grid[m-1][n-1]
}
