package search_range

func searchRange(nums []int, target int) []int {
    if len(nums) == 0 {
		return []int{-1, -1}
	}
	var min, max = 0, len(nums) - 1
	var half int
	for min < max {
		half = (min + max) / 2
		if nums[min] < nums[half] {
			if nums[(min + half) / 2 + 1] < target {
				min = (min + half) / 2 + 1
			} else {
				min ++
			}
			continue
		}
		if nums[max] > nums[half] {
			if nums[(max + half) / 2] > target {
				max = (max + half) / 2
			} else {
				max --
			}
			continue
		}
		return []int{min, max}
	}
	if min == max && nums[min] == target {
		return []int{min, min}
	}
	return []int{-1, -1}
}

func searchRangeV1(nums []int, target int) []int {
    if len(nums) == 0 {
		return []int{-1, -1}
	}
	var min, max = 0, len(nums) - 1
	for min < max {
		half := (min + max) / 2
		if nums[half] < target {
			min = half + 1
		} else if nums[half] > target {
			max = half
		} else {
			var start, end = half, half
			for {
				if start > 0 && nums[start - 1] == target {
					start --
					continue
				}
				if end < len(nums) - 1 && nums[end + 1] == target {
					end ++
					continue
				}
				return []int{start, end}
			}
		}
	}
	if min == max && nums[min] == target {
		return []int{min, min}
	}
	return []int{-1, -1}
}