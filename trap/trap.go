// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。



// 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。

// 示例:

// 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出: 6

package trap

func trap(height []int) int {
	ret := 0
	l := len(height)
	if l <= 2 {
		return ret
	}

	s := make([]int, 0)

	for i := range height {
		if len(s) == 0 {
			s = append(s, i)
			continue
		}

		idx := s[len(s)-1]
		if height[i] < height[idx] {
			s = append(s, i)
			continue
		}

		for len(s) > 0 {
			idx := s[len(s)-1]
			area := (i-idx)*
		}
	}
}

// func trap(height []int) int {
// 	if len(height) < 2 {
// 		return 0
// 	}
// 	left := 0
// 	area, ret := 0, 0
// 	for i := 1; i < len(height); i++ {
// 		if height[i] < height[left] {
// 			area += height[left] - height[i]
// 			continue
// 		} else {
// 			ret += area
// 			area = 0
// 			left = i
// 		}
// 	}
// 	if left < len(height)-2 {
// 		right := len(height) - 1
// 		area = 0
// 		for i := right - 1; i >= left; i-- {
// 			if height[i] < height[right] {
// 				area += height[right] - height[i]
// 			} else {
// 				ret += area
// 				area = 0
// 				right = i
// 			}
// 		}
// 	}
// 	return ret
// }
