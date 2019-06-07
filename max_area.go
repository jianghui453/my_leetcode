package leet_code

import (
	"math"
)

func maxArea(height []int) int {
	var left, right, area int
	left = 0
	right = len(height) - 1
	for left < right{
		if int(math.Min(float64(height[left]), float64(height[right]))) * (right - left) > area {
			area = int(math.Min(float64(height[left]), float64(height[right]))) * (right - left)
		}
		if height[left] > height[right] {
			right --
		} else {
			left ++
		}
	}
	return area
}