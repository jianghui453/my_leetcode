package array

import "testing"

func TestSpiralOrder(t *testing.T) {
	var matrix [][]int
	var h, r []int

	matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	h = []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	r = spiralOrder(matrix)
	t.Logf("\nh=%v \nr=%v", h, r)

	matrix = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	h = []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}
	r = spiralOrder(matrix)
	t.Logf("\nh=%v \nr=%v", h, r)

	matrix = [][]int{
		{7},
		{9},
		{6},
	}
	h = []int{7, 9, 6}
	r = spiralOrder(matrix)
	t.Logf("\nh=%v \nr=%v", h, r)
}
