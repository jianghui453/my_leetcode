package array

import "testing"

func TestMinimumTotal(t *testing.T) {
	var triangle [][]int
	var h, r int

	triangle = [][]int{
		{2},
		{3,4},
		{6,5,7},
		{4,1,8,3},
	}
	h = 11
	r = minimumTotal(triangle)
	t.Logf("%t h=%d r=%d", h==r, h, r)
}
