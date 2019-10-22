package merge

import "testing"

func TestMerge(t *testing.T) {
	var intervals [][]int
	var ret, hope [][]int

	intervals = [][]int{
		{1, 3}, {2, 6}, {8, 10}, {15, 18},
	}
	hope = [][]int{
		{1, 6}, {8, 10}, {15, 18},
	}
	ret = merge(intervals)
	t.Logf("intervals=%v \nhope=%v \nret=%v", intervals, hope, ret)

	intervals = [][]int{
		{1, 4}, {4, 5},
	}
	hope = [][]int{
		{1, 5},
	}
	ret = merge(intervals)
	t.Logf("intervals=%v \nhope=%v \nret=%v", intervals, hope, ret)

	intervals = [][]int{
		{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10},
	}
	hope = [][]int{
		{1, 10},
	}
	ret = merge(intervals)
	t.Logf("intervals=%v \nhope=%v \nret=%v", intervals, hope, ret)

	intervals = [][]int{
		{2, 3}, {6, 7}, {4, 5}, {8, 9}, {1, 10},
	}
	hope = [][]int{
		{1, 10},
	}
	ret = merge(intervals)
	t.Logf("intervals=%v \nhope=%v \nret=%v", intervals, hope, ret)
}
