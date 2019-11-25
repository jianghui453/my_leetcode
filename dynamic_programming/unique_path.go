//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
//
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
//
//问总共有多少条不同的路径？
//
//
//
//例如，上图是一个7 x 3 的网格。有多少可能的路径？
//
//说明：m 和 n 的值均不超过 100。
//
//示例 1:
//
//输入: m = 3, n = 2
//输出: 3
//解释:
//从左上角开始，总共有 3 条路径可以到达右下角。
//1. 向右 -> 向右 -> 向下
//2. 向右 -> 向下 -> 向右
//3. 向下 -> 向右 -> 向右
//示例 2:
//
//输入: m = 7, n = 3
//输出: 28

package dynamic_programming

func uniquePaths(m int, n int) int {
	if m*n == 0 {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j > 0 && i > 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			} else if j == 0 && i > 0 {
				dp[i][j] = dp[i-1][j]
			} else if j > 0 && i == 0 {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

//func uniquePaths(m int, n int) int {
//	if m < 2 && n < 2 {
//		return 0
//	}
//	ret := 0
//	_uniquePaths(m, n, 1, 1, &ret)
//	return ret
//}
//
//func _uniquePaths(m, n, _m, _n int, ret *int) {
//	if _m >= m && _n >= n {
//		*ret++
//		return
//	}
//	if _m < m {
//		_uniquePaths(m, n, _m+1, _n, ret)
//	}
//	if _n < n {
//		_uniquePaths(m, n, _m, _n+1, ret)
//	}
//}
