package remove_element

func removeElement(nums []int, val int) int {
	count, i := 0, 0
	for i < len(nums) - count {
		if nums[i] == val {
			nums = append(nums[:i], append(nums[i+1:], nums[i])...)
			count ++
			continue
		}
		i ++
	}
	return len(nums) - count
}