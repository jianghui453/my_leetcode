package trap

func trap(height []int) int {
	if len(height) < 2 {
		return 0
	}
	left := 0
	area, ret := 0, 0
	for i := 1; i < len(height); i++ {
		if height[i] < height[left] {
			area += height[left] - height[i]
			continue
		} else {
			ret += area
			area = 0
			left = i
		}
	}
	if left < len(height)-2 {
		right := len(height) - 1
		area = 0
		for i := right - 1; i >= left; i-- {
			if height[i] < height[right] {
				area += height[right] - height[i]
			} else {
				ret += area
				area = 0
				right = i
			}
		}
	}
	return ret
}
