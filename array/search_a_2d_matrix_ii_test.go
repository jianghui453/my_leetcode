package array

import (
	"testing"
)

func Test_searchMatrix(t *testing.T) {
	var matrix [][]int
	var target int
	var hope, ret bool

	matrix = [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	target = 5
	hope = true
	ret = searchMatrix(matrix, target)
	t.Log(ret == hope, matrix, hope, ret)

	target = 20
	hope = false
	ret = searchMatrix(matrix, target)
	t.Log(ret == hope, matrix, hope, ret)
}
