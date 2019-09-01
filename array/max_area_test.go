package array

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	var height []int
	var hope, ret int

	height = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	hope = 49
	ret = maxArea(height)
	if hope == ret {
		t.Logf("success: height = %v; hope = %d; ret = %d\n", height, hope, ret)
	} else {
		t.Errorf("fail: height = %v; hope = %d; ret = %d\n", height, hope, ret)
	}
}
