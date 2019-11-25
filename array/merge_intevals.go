// Given a collection of intervals, merge all overlapping intervals.

// Example 1:

// Input: [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
// Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
// Example 2:

// Input: [[1,4],[4,5]]
// Output: [[1,5]]
// Explanation: Intervals [1,4] and [4,5] are considered overlapping.
// NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.

package array

func merge(intervals [][]int) [][]int {
	l := len(intervals)
	if l <= 1 {
		return intervals
	}

	ret := make([][]int, 0)
	ret = append(ret, intervals[0])
	rl := 1

loop:
	for i := 0; i < l; i++ {
		for j := 0; j < rl; j++ {
			if intervals[i][0] < ret[j][0] && intervals[i][1] > ret[j][1] {
				copy(ret[j], intervals[i])
				continue loop
			}

			if intervals[i][0] >= ret[j][0] && intervals[i][0] <= ret[j][1] {
				if intervals[i][1] > ret[j][1] {
					ret[j][1] = intervals[i][1]
				}
				continue loop
			}

			if intervals[i][1] >= ret[j][0] && intervals[i][1] <= ret[j][1] {
				if intervals[i][0] < ret[j][0] {
					ret[j][0] = intervals[i][0]
				}
				continue loop
			}
		}

		ret = append(ret, intervals[i])
		rl++
	}

	if l == rl {
		return ret
	}

	return merge(ret)
}

// func merge(intervals [][]int) [][]int {
// 	ret := make([][]int, 0)
// 	l := len(intervals)
// 	if l < 1 {
// 		return ret
// 	}

// 	quickSort(intervals)

// 	ret = append(ret, intervals[0])
// 	for i := 1; i < l; i++ {
// 		if intervals[i][0] > ret[len(ret)-1][1] {
// 			ret = append(ret, intervals[i])
// 		} else if intervals[i][1] > ret[len(ret)-1][1] {
// 			ret[len(ret)-1][1] = intervals[i][1]
// 		}
// 	}

// 	return ret
// }

// func quickSort(intervals [][]int) {
// 	l := len(intervals)
// 	if l <= 1 {
// 		return
// 	}

// 	newItv := make([][]int, l)

// 	val := intervals[0][0]
// 	left, right := 0, l-1
// 	for i := 1; i < l; i++ {
// 		if intervals[i][0] >= val {
// 			newItv[right] = make([]int, 2)
// 			copy(newItv[right], intervals[i])
// 			right--
// 		} else if intervals[i][0] < val {
// 			newItv[left] = make([]int, 2)
// 			copy(newItv[left], intervals[i])
// 			left++
// 		}
// 	}

// 	newItv[right] = make([]int, 2)
// 	copy(newItv[right], intervals[0])

// 	if left < l {
// 		quickSort(newItv[: left])
// 	}

// 	if right < l-1 {
// 		quickSort(newItv[right+1: ])
// 	}

// 	copy(intervals, newItv)
// }

// func merge(intervals [][]int) [][]int {
// 	iLen := len(intervals)
// 	if iLen > 1 {
// 		for {
// 			intervals = _merge(intervals)
// 			_iLen := len(intervals)
// 			if _iLen == iLen {
// 				break
// 			}
// 			iLen = _iLen
// 		}

// 	}
// 	return intervals
// }

// func _merge(intervals [][]int) [][]int {
// 	var ret [][]int
// 	iLen := len(intervals)
// Loop:
// 	for i := 0; i < iLen; i++ {
// 		for j := 0; j < len(ret); j++ {
// 			if ret[j][1] < intervals[i][0] || ret[j][0] > intervals[i][1] {
// 				continue
// 			}
// 			var min, max int
// 			if ret[j][0] < intervals[i][0] {
// 				min = ret[j][0]
// 			} else {
// 				min = intervals[i][0]
// 			}
// 			if ret[j][1] > intervals[i][1] {
// 				max = ret[j][1]
// 			} else {
// 				max = intervals[i][1]
// 			}
// 			ret[j] = []int{min, max}
// 			continue Loop
// 		}
// 		ret = append(ret, intervals[i])
// 	}
// 	return ret
// }
