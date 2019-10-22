package matrix

import "testing"

func TestSearchMatrix(t *testing.T) {
	var matrix [][]int
	var hope, ret bool
	var target int

	matrix = [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	target = 3
	hope = true
	ret = searchMatrix(matrix, target)
	t.Logf("matrix=%v hope=%t ret=%t", matrix, hope, ret)

	matrix = [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	target = 12
	hope = false
	ret = searchMatrix(matrix, target)
	t.Logf("matrix=%v hope=%t ret=%t", matrix, hope, ret)
}
