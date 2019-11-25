package binary_search

import "testing"

func TestSearchMatrix(t *testing.T) {
	var matrix [][]int
	var hope, ret bool
	var target int

	matrix = [][]int{{1}}
	target = 0
	hope = false
	ret = searchMatrix(matrix, target)
	t.Logf("%t matrix=%v target=%d hope=%t ret=%t", ret == hope, matrix, target, hope, ret)

	matrix = [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	target = 3
	hope = true
	ret = searchMatrix(matrix, target)
	t.Logf("%t matrix=%v target=%d hope=%t ret=%t", ret == hope, matrix, target, hope, ret)

	matrix = [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	target = 12
	hope = false
	ret = searchMatrix(matrix, target)
	t.Logf("%t matrix=%v target=%d hope=%t ret=%t", ret == hope, matrix, target, hope, ret)

	matrix = [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}, {51, 52, 53, 54}}
	target = 51
	hope = true
	ret = searchMatrix(matrix, target)
	t.Logf("%t matrix=%v target=%d hope=%t ret=%t", ret == hope, matrix, target, hope, ret)

	matrix = [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}, {51, 52, 53, 55}}
	target = 54
	hope = false
	ret = searchMatrix(matrix, target)
	t.Logf("%t matrix=%v target=%d hope=%t ret=%t", ret == hope, matrix, target, hope, ret)
}
