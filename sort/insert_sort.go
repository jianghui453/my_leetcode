package sort

func insertSort(nums []int) {
	nLen := len(nums)
	for i := 1; i < nLen; i++ {
		num := nums[i]
		for j := i - 1; j >= 0; j-- {
			if nums[j] > num {
				nums[j+1] = nums[j]
			} else {
				nums[j+1] = num
				break
			}
		}
		if nums[0] > num {
			nums[0] = num
		}
	}
}
