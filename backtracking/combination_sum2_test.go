package backtracking

import (
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	var candinates []int
	var target int
	var h, r [][]int

	candinates = []int{10, 1, 2, 7, 6, 1, 5}
	target = 8
	h = [][]int{
		{1, 7},
		{1, 2, 5},
		{2, 6},
		{1, 1, 6},
	}
	r = combinationSum2(candinates, target)
	t.Logf("candinates=%v \ntarget=%d \nh=%v \nr=%v", candinates, target, h, r)

	candinates = []int{2, 5, 2, 1, 2}
	target = 5
	h = [][]int{
		{1, 2, 2},
		{5},
	}
	r = combinationSum2(candinates, target)
	t.Logf("candinates=%v \ntarget=%d \nh=%v \nr=%v", candinates, target, h, r)
}
