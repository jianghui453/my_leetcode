package sort

import "testing"

func TestInsertSort(t *testing.T) {
	var nums []int

	nums = []int{8, 5, 2, 6, 9, 3, 1, 4, 0, 7}
	t.Logf("start nums=%v", nums)
	insertSort(nums)
	t.Logf("end nums=%v", nums)
}
