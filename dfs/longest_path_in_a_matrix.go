// 给定一个整数矩阵，找出最长递增路径的长度。

// 对于每个单元格，你可以往上，下，左，右四个方向移动。 你不能在对角线方向上移动或移动到边界外（即不允许环绕）。

// 示例 1:

// 输入: nums =
// [
//   [9,9,4],
//   [6,6,8],
//   [2,1,1]
// ]
// 输出: 4
// 解释: 最长递增路径为 [1, 2, 6, 9]。
// 示例 2:

// 输入: nums =
// [
//   [3,4,5],
//   [3,2,6],
//   [2,2,1]
// ]
// 输出: 4
// 解释: 最长递增路径是 [3, 4, 5, 6]。注意不允许在对角线方向上移动。

package dfs

func longestIncreasingPath(matrix [][]int) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}

	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
	}

	var ret int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ret = max(ret, dfs(matrix, &memo, m, n, i, j))
		}
	}

	return ret
}

func dfs(matrix [][]int, memo *[][]int, m, n, x, y int) int {
	if (*memo)[x][y] == 0 {
		(*memo)[x][y] = 1

		if x < m-1 && matrix[x][y] < matrix[x+1][y] {
			(*memo)[x][y] = dfs(matrix, memo, m, n, x+1, y) + 1
		}

		if x > 0 && matrix[x][y] < matrix[x-1][y] {
			(*memo)[x][y] = max((*memo)[x][y], dfs(matrix, memo, m, n, x-1, y)+1)
		}

		if y < n-1 && matrix[x][y] < matrix[x][y+1] {
			(*memo)[x][y] = max((*memo)[x][y], dfs(matrix, memo, m, n, x, y+1)+1)
		}

		if y > 0 && matrix[x][y] < matrix[x][y-1] {
			(*memo)[x][y] = max((*memo)[x][y], dfs(matrix, memo, m, n, x, y-1)+1)
		}
	}

	return (*memo)[x][y]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
