package sort

import "testing"

func TestRadioSort(t *testing.T) {
    var nums []int

    nums = []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
    //nums = []int{2, 1}
    t.Logf("start nums=%v", nums)
    radioSort(nums)
    t.Logf("end nums=%v", nums)
}
