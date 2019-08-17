package sort

import (
	"testing"
)

func TestShellSort(t *testing.T) {
	var nums []int

	nums = []int{8, 5, 2, 6, 9, 3, 1, 4, 0, 7}
	t.Logf("start nums=%v", nums)
	shellSort(nums)
	t.Logf("end nums=%v", nums)

	//var nums1, nums2 []int
	//for i := 0; i <= 10; i ++ {
	//     num := rand.Int()
	//     nums1 = append(nums1, num)
	//     nums2 = append(nums2, num)
	//}
	//shellSort(nums1)
	//insertSort(nums2)
}
