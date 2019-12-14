package two_pointer

import (
	"testing"
)

func Test_twoSum(t *testing.T) {
	var numbers, hope, ret []int
	var target int

	numbers = []int{2, 7, 11, 15}
	target = 9
	hope = []int{1, 2}
	ret = twoSum(numbers, target)
	t.Log(numbers, target, hope, ret)
}