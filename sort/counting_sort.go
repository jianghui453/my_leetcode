package sort

func countingSort(nums []int) {
	l := len(nums)

	max := 0
	for i := 0; i < l; i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	bucket := make([]int, max+1)
	for i := 0; i < l; i++ {
		bucket[nums[i]]++
	}

	j := 0
	for i := 0; i <= max; i++ {
		for ; bucket[i] > 0; bucket[i]-- {
			nums[j] = i
			j++
		}
	}
}
