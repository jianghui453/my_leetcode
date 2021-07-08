package sort

import (
	"testing"
)

func TestShellSort(t *testing.T) {
	var nums []int

	nums = []int{8, 5, 2, 6, 9, 3, 1, 4, 0, 7}
	t.Logf("start nums=%v", nums)
	shellSort(nums)
	t.Logf("  end nums=%v", nums)
}
