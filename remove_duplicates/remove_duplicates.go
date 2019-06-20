package remove_duplicates

func removeDuplicates(nums []int) int {
	count := 0
	i := 0
	rptMap := make(map[int]bool)
	for i < len(nums)-count {
		if _, ok := rptMap[nums[i]]; ok {
			nums = append(nums[:i], append(nums[i+1:], nums[i])...)
			count++
			continue
		}
		rptMap[nums[i]] = true
		i++
	}
	return i
}
