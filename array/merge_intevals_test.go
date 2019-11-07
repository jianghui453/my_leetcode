package array

import "testing"

func TestMerge(t *testing.T) {
	var intervals [][]int
	var r, h [][]int

	intervals = [][]int{
		{1, 3}, {2, 6}, {8, 10}, {15, 18},
	}
	h = [][]int{
		{1, 6}, {8, 10}, {15, 18},
	}
	r = merge(intervals)
	t.Logf("intervals=%v \nh=%v \nr=%v", intervals, h, r)

	intervals = [][]int{
		{1, 4}, {4, 5},
	}
	h = [][]int{
		{1, 5},
	}
	r = merge(intervals)
	t.Logf("intervals=%v \nh=%v \nr=%v", intervals, h, r)

	intervals = [][]int{
		{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10},
	}
	h = [][]int{
		{1, 10},
	}
	r = merge(intervals)
	t.Logf("intervals=%v \nh=%v \nr=%v", intervals, h, r)

	intervals = [][]int{
		{2, 3}, {6, 7}, {4, 5}, {8, 9}, {1, 10},
	}
	h = [][]int{
		{1, 10},
	}
	r = merge(intervals)
	t.Logf("intervals=%v \nh=%v \nr=%v", intervals, h, r)

	intervals = [][]int{
		{2, 3},
	}
	h = [][]int{
		{2, 3},
	}
	r = merge(intervals)
	t.Logf("intervals=%v \nh=%v \nr=%v", intervals, h, r)

	intervals = [][]int{
		{1, 4}, {1, 5},
	}
	h = [][]int{
		{1, 5},
	}
	r = merge(intervals)
	t.Logf("intervals=%v \nh=%v \nr=%v", intervals, h, r)

	intervals = [][]int{
		{1, 3}, {4, 5},
	}
	h = [][]int{
		{1, 3}, {4, 5},
	}
	r = merge(intervals)
	t.Logf("intervals=%v \nh=%v \nr=%v", intervals, h, r)
}
