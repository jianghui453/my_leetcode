package array

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	var height []int
	var h, r int

	height = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	h = 49
	r = maxArea(height)
	t.Logf("%t height=%v h=%d r=%d", r == h, height, h, r)
}
