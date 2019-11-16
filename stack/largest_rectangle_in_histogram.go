//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
//求在该柱状图中，能够勾勒出来的矩形的最大面积。
//
//
//
//
//
//以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。
//
//
//
//
//
//图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。
//
//
//
//示例:
//
//输入: [2,1,5,6,2,3]
//输出: 10

package stack

import (
	// "fmt"
	"math"
)

func largestRectangleArea(heights []int) int {
	l := len(heights)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return heights[0]
	}

	area := 0
	s := []int{-1}
	for i := 0; i < l; i++ {
		for len(s) > 1 && heights[i] < heights[s[len(s)-1]] {
			area = int(math.Max(float64(heights[s[len(s)-1]] * (i - 1 - s[len(s)-2])), float64(area)))
			s = s[: len(s)-1]
		}
		s = append(s, i)
	}

	ls := len(s)
	for i := ls-1; i > 0; i-- {
		area = int(math.Max(float64(heights[s[i]] * (l - 1 - s[i-1])), float64(area)))
	}

	return area
}

// func largestRectangleArea(heights []int) int {
// 	lenH := len(heights)
// 	if lenH < 1 {
// 		return 0
// 	}
// 	if lenH < 2 {
// 		return heights[0]
// 	}
// 	s := []int{-1}
// 	lenS := 1
// 	area := 0
// 	for i := 0; i < lenH; i++ {
// 		for lenS > 1 && heights[i] < heights[s[lenS-1]] {
// 			left := s[lenS-2]
// 			right := i - 1
// 			width := right - left
// 			height := heights[s[lenS-1]]
// 			areaI := width * height
// 			if areaI > area {
// 				area = areaI
// 			}
// 			s = s[:lenS-1]
// 			lenS--
// 		}
// 		s = append(s, i)
// 		lenS++
// 	}
// 	fmt.Printf("s=%v lenS=%d lenH=%d\n", s, lenS, lenH)
// 	idxS := lenS - 1
// 	idxH := lenH
// 	for ; idxS > 0; idxS-- {
// 		left := s[idxS-1]
// 		right := idxH - 1
// 		width := right - left
// 		height := heights[s[idxS]]
// 		areaI := width * height
// 		if areaI > area {
// 			area = areaI
// 		}
// 	}
// 	return area
// }

//func largestRectangleArea(heights []int) int {
//    lenH := len(heights)
//    if lenH < 1 {
//        return 0
//    }
//    if lenH < 2 {
//        return heights[0]
//    }
//    area := heights[0]
//    for i := 0; i < lenH; i ++ {
//        left := i
//        right := i
//        for {
//            val := 0
//            if left > 0 && heights[left-1] >= heights[i] {
//                left --
//                val ++
//            }
//            if right < lenH-1 && heights[right+1] >= heights[i] {
//                right ++
//                val ++
//            }
//            if val < 1 {
//                break
//            }
//        }
//        areaI := (right-left+1)*heights[i]
//        if areaI > area {
//            area = areaI
//        }
//    }
//    return area
//}
