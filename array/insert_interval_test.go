package array

import "testing"

func TestInsert(t *testing.T) {
	var intervals, h, r [][]int
	var newInterval []int

	intervals = [][]int{
		{1, 3}, {6, 9},
	}
	newInterval = []int{2, 5}
	h = [][]int{{1, 5}, {6, 9}}
	r = insert(intervals, newInterval)
	t.Logf("\nintervals=%v \nnewInterval=%v \nh=%v \nr=%v", intervals, newInterval, h, r)

	intervals = [][]int{
		{1, 3}, {6, 9},
	}
	newInterval = []int{1, 10}
	h = [][]int{{1, 10}}
	r = insert(intervals, newInterval)
	t.Logf("\nintervals=%v \nnewInterval=%v \nh=%v \nr=%v", intervals, newInterval, h, r)

	intervals = [][]int{
		{1, 3},
	}
	newInterval = []int{2, 5}
	h = [][]int{{1, 5}}
	r = insert(intervals, newInterval)
	t.Logf("\nintervals=%v \nnewInterval=%v \nh=%v \nr=%v", intervals, newInterval, h, r)

	intervals = [][]int{
		{6, 9},
	}
	newInterval = []int{2, 5}
	h = [][]int{{2, 5}, {6, 9}}
	r = insert(intervals, newInterval)
	t.Logf("\nintervals=%v \nnewInterval=%v \nh=%v \nr=%v", intervals, newInterval, h, r)

	intervals = [][]int{
		{6, 9},
	}
	newInterval = []int{2, 6}
	h = [][]int{{2, 9}}
	r = insert(intervals, newInterval)
	t.Logf("\nintervals=%v \nnewInterval=%v \nh=%v \nr=%v", intervals, newInterval, h, r)

	intervals = [][]int{
		{1, 3}, {6, 9},
	}
	newInterval = []int{4, 5}
	h = [][]int{{1, 3}, {4, 5}, {6, 9}}
	r = insert(intervals, newInterval)
	t.Logf("\nintervals=%v \nnewInterval=%v \nh=%v \nr=%v", intervals, newInterval, h, r)

	intervals = [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newInterval = []int{4, 8}
	h = [][]int{{1, 2}, {3, 10}, {12, 16}}
	r = insert(intervals, newInterval)
	t.Logf("\nintervals=%v \nnewInterval=%v \nh=%v \nr=%v", intervals, newInterval, h, r)

	intervals = [][]int{}
	newInterval = []int{4, 8}
	h = [][]int{{4, 8}}
	r = insert(intervals, newInterval)
	t.Logf("\nintervals=%v \nnewInterval=%v \nh=%v \nr=%v", intervals, newInterval, h, r)
}
