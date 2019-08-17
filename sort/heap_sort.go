package sort

import "fmt"

func heapSort(nums []int) {
	nLen := len(nums)
	if nLen < 2 {
		return
	}
	for i := nLen/2-1; i >= 0; i-- {
		heapify(nums, nLen, i)
	}
    //heapify(nums, nLen, 0)
    fmt.Printf("nums=%v\n", nums)
	nums[nLen-1], nums[0] = nums[0], nums[nLen-1]
	for i := nLen-2; i >= 0; i-- {
        fmt.Printf("nums[:i+1]=%v\n", nums[:i+1])
		heapify(nums[:i+1], i+1, 0)
		nums[i], nums[0] = nums[0], nums[i]
	}
}

// heapify confirm
func heapify(nums []int, n, i int) {
	//fmt.Printf("nums=%v n=%d i=%d\n", nums, n, i)
	left := 2*i + 1
	right := 2*i + 2
	if right < n && nums[right] > nums[i] {
        nums[right], nums[i] = nums[i], nums[right]
        heapify(nums, n, right)
	}
	if left < n && nums[left] > nums[i] {
        nums[left], nums[i] = nums[i], nums[left]
        heapify(nums, n, left)
	}
}
