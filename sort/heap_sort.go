package sort

import (
	// "fmt"
)

func heapSort(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		heapify(nums, len(nums)-1, i)
	}

	for i := len(nums)-1; i >= 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heapify(nums, i, 0)
	}
}

// heapify confirm
func heapify(nums []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && nums[left] > nums[largest] {
		largest = left
	}

	if right < n && nums[right] > nums[largest] {
		largest = right
	}

	if largest != i {
		nums[largest], nums[i] = nums[i], nums[largest]
		heapify(nums, n, largest)
	}
}
