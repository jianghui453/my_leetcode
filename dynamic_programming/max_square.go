// 在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。

// 示例:

// 输入: 

// 1 0 1 0 0
// 1 0 1 1 1
// 1 1 1 1 1
// 1 0 0 1 0

// 输出: 4

package dynamic_programming

import (
	// "fmt"
)

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}

	n := len(matrix[0])
	if n == 0 {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	ret := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}

			dp[i][j] = 1
			if i > 0 && j > 0 {
				for k := 1; k <= i && k <= j && k <= dp[i-1][j-1]; k++ {
					if matrix[i-k][j] == '1' && matrix[i][j-k] == '1' {
						dp[i][j] = k+1
					} else {
						break
					}
				}
			}
			
			ret = max(ret, dp[i][j] * dp[i][j])
		}
	}
	
	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}