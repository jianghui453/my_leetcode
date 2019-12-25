package dfs

import (
	"testing"
)

func Test_longestIncreasingPath(t *testing.T) {
	var matrix [][]int
	var hope, ret int

	matrix = [][]int{
		{3, 4, 5},
		{3, 2, 6},
		{2, 2, 1},
	}
	hope = 4
	ret = longestIncreasingPath(matrix)
	t.Log(ret == hope, matrix, hope, ret)

	matrix = [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}
	hope = 4
	ret = longestIncreasingPath(matrix)
	t.Log(ret == hope, matrix, hope, ret)
}
