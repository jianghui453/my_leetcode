package sort

import (
	"testing"
)

func TestHeapSort(t *testing.T) {
	var nums []int

	nums = []int{8, 5, 2, 6, 9, 3, 1, 4, 0, 7}
	t.Log("before =", nums)
	heapSort(nums)
	t.Log("after =", nums)
}
