package sort

func quickSort(nums []int) {
	l := len(nums)
	if l <= 1 {
		return
	}

	newnums := make([]int, l)
	left, right := 0, l-1
	for i := 1; i < l; i++ {
		if nums[i] > nums[0] {
			newnums[right], right = nums[i], right-1
		} else if nums[i] < nums[0] {
			newnums[left], left = nums[i], left+1
		}
	}
	for i := left; i <= right; i++ {
		newnums[i] = nums[0]
	}
	
	copy(nums, newnums)
	quickSort(nums[: left])
	if right < l-1 {
		quickSort(nums[right+1: ])
	}
}
