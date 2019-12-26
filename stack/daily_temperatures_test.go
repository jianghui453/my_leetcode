package stack

import (
	"testing"
)

func Test_dailyTemperatures(t *testing.T) {
	var (
		T, hope, ret []int
	)

	T = []int{73, 74, 75, 71, 69, 72, 76, 73}
	hope = []int{1, 1, 4, 2, 1, 1, 0, 0}
	ret = dailyTemperatures(T)
	t.Log(T)
	t.Log(hope)
	t.Log(ret)

	T = []int{73, 74, 75, 71, 69, 72, 76, 77}
	hope = []int{1, 1, 4, 2, 1, 1, 1, 0}
	ret = dailyTemperatures(T)
	t.Log(T)
	t.Log(hope)
	t.Log(ret)

	T = []int{60, 59, 58, 57, 56}
	hope = []int{0, 0, 0, 0, 0}
	ret = dailyTemperatures(T)
	t.Log(T)
	t.Log(hope)
	t.Log(ret)

	T = []int{50, 51, 52, 53, 54, 55}
	hope = []int{1, 1, 1, 1, 1, 0}
	ret = dailyTemperatures(T)
	t.Log(T)
	t.Log(hope)
	t.Log(ret)
}
