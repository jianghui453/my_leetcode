package dynamic_programming

import "testing"

func TestMinPathSum(t *testing.T) {
	var grid [][]int
	var h, r int

	grid = [][]int{
		{1,3,1},
		{1,5,1},
		{4,2,1},
	}
	h = 7
	r = minPathSum(grid)
	t.Logf("h=%d r=%d", h, r)

	grid = [][]int{
		{1},
	}
	h = 1
	r = minPathSum(grid)
	t.Logf("h=%d r=%d", h, r)

	grid = [][]int{
		{1, 2, 5},
		{3, 2, 1},
	}
	h = 6
	r = minPathSum(grid)
	t.Logf("h=%d r=%d", h, r)
}