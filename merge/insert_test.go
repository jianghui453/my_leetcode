package merge

import "testing"

func TestInsert(t *testing.T) {
	var intervals, hope, ret [][]int
	var newInterval []int

	intervals = [][]int{
		{1, 3}, {6, 9},
	}
	newInterval = []int{2, 5}
	hope = [][]int{{1, 5}, {6, 9}}
	ret = insert(intervals, newInterval)
	t.Logf("\nintervals=%v \nnewInterval=%v \nhope=%v \nret=%v", intervals, newInterval, hope, ret)

	intervals = [][]int{
		{1, 3}, {6, 9},
	}
	newInterval = []int{1, 10}
	hope = [][]int{{1, 10}}
	ret = insert(intervals, newInterval)
	t.Logf("\nintervals=%v \nnewInterval=%v \nhope=%v \nret=%v", intervals, newInterval, hope, ret)

	intervals = [][]int{
		{1, 3},
	}
	newInterval = []int{2, 5}
	hope = [][]int{{1, 5}}
	ret = insert(intervals, newInterval)
	t.Logf("\nintervals=%v \nnewInterval=%v \nhope=%v \nret=%v", intervals, newInterval, hope, ret)

	intervals = [][]int{
		{6, 9},
	}
	newInterval = []int{2, 5}
	hope = [][]int{{2, 5}, {6, 9}}
	ret = insert(intervals, newInterval)
	t.Logf("\nintervals=%v \nnewInterval=%v \nhope=%v \nret=%v", intervals, newInterval, hope, ret)

	intervals = [][]int{
		{6, 9},
	}
	newInterval = []int{2, 6}
	hope = [][]int{{2, 9}}
	ret = insert(intervals, newInterval)
	t.Logf("\nintervals=%v \nnewInterval=%v \nhope=%v \nret=%v", intervals, newInterval, hope, ret)

	intervals = [][]int{
		{1, 3}, {6, 9},
	}
	newInterval = []int{4, 5}
	hope = [][]int{{1, 3}, {4, 5}, {6, 9}}
	ret = insert(intervals, newInterval)
	t.Logf("\nintervals=%v \nnewInterval=%v \nhope=%v \nret=%v", intervals, newInterval, hope, ret)

	intervals = [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newInterval = []int{4, 8}
	hope = [][]int{{1, 2}, {3, 10}, {12, 16}}
	ret = insert(intervals, newInterval)
	t.Logf("\nintervals=%v \nnewInterval=%v \nhope=%v \nret=%v", intervals, newInterval, hope, ret)

	intervals = [][]int{}
	newInterval = []int{4, 8}
	hope = [][]int{{4, 8}}
	ret = insert(intervals, newInterval)
	t.Logf("\nintervals=%v \nnewInterval=%v \nhope=%v \nret=%v", intervals, newInterval, hope, ret)
}
