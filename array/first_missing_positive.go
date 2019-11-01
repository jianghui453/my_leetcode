package array

func firstMissingPositive(nums []int) int {
	hash := make(map[int]bool)
	max := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			continue
		}
		if nums[i] > max {
			max = nums[i]
		}
		if _, ok := hash[nums[i]]; !ok {
			hash[nums[i]] = true
		}
	}
	if max == 0 {
		return 1
	}

	var i int
	for i = 1; i < max; i++ {
		if _, ok := hash[i]; !ok {
			return i
		}
	}
	return i + 1
}
