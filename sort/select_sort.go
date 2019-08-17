package sort

func selectSort(nums []int) {
	nLen := len(nums)
	for i := 0; i < nLen-1; i++ {
		minIdx := i
		minVal := nums[i]
		for j := i + 1; j < nLen; j++ {
			if nums[j] < minVal {
				minVal = nums[j]
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
}
