// Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).

// You may assume that the intervals were initially sorted according to their start times.

// Example 1:

// Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
// Output: [[1,5],[6,9]]
// Example 2:

// Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
// Output: [[1,2],[3,10],[12,16]]
// Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
// NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.

package array

import (
	"math"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	l := len(intervals)
	ret := make([][]int, 0)
	if l == 0 {
		ret = append(ret, newInterval)
		return ret
	}

	if newInterval[1] < intervals[0][0] {
		ret = append(append(ret, newInterval), intervals...)
		return ret
	}

	if newInterval[0] > intervals[l-1][1] {
		ret = append(intervals, newInterval)
		return ret
	}

	for i := 0; i < l; i++ {
		if newInterval[0] > intervals[i][1] {
			continue
		}

		if newInterval[1] < intervals[i][0] {
			ret = append(ret, intervals[: i]...)
			ret = append(ret, newInterval)
			ret = append(ret, intervals[i: ]...)
			return ret
		}

		
		for j := i; j < l; j++ {
			if newInterval[1] < intervals[j][0] {
				ret = append(ret, intervals[: i]...)
				ret = append(ret, []int{int(math.Min(float64(newInterval[0]), float64(intervals[i][0]))), newInterval[1]})
				ret = append(ret, intervals[j: ]...)
				return ret
			}

			if newInterval[1] <= intervals[j][1] {
				ret = append(ret, intervals[: i]...)
				ret = append(ret, []int{int(math.Min(float64(newInterval[0]), float64(intervals[i][0]))), intervals[j][1]})
				if j < l-1 {
					ret = append(ret, intervals[j+1: ]...)
				}
				return ret
			}
		}
		
		ret = append(ret, intervals[: i]...)
		ret = append(ret, []int{int(math.Min(float64(newInterval[0]), float64(intervals[i][0]))), newInterval[1]})
		return ret
	}

	return ret
}

// func insert(intervals [][]int, newInterval []int) [][]int {
// 	var ret [][]int
// 	iLen := len(intervals)
// 	if iLen == 0 {
// 		return append(ret, newInterval)
// 	}
// 	if newInterval[1] < intervals[0][0] {
// 		ret = append(ret, newInterval)
// 		ret = append(ret, intervals...)
// 		return ret
// 	}
// 	if newInterval[0] > intervals[iLen-1][1] {
// 		ret = append(ret, intervals...)
// 		ret = append(ret, newInterval)
// 		return ret
// 	}
// 	var head, tail [][]int
// 	for i := 0; i < iLen; i++ {
// 		if newInterval[0] > intervals[i][0] {
// 			head = append(head, intervals[i])
// 			continue
// 		}
// 		if i == 0 {
// 			head = append(head, newInterval)
// 			break
// 		}
// 		if newInterval[0] > intervals[i-1][1] {
// 			head = append(head, newInterval)
// 			break
// 		}
// 		if newInterval[1] > intervals[i-1][1] {
// 			head[len(head)-1][1] = newInterval[1]
// 		}
// 	}
// 	for i := iLen - 1; i >= 0; i-- {
// 		if intervals[i][0] > newInterval[1] {
// 			tail = append([][]int{{intervals[i][0], intervals[i][1]}}, tail...)
// 			continue
// 		}
// 		if intervals[i][1] > newInterval[1] {
// 			head[len(head)-1][1] = intervals[i][1]
// 		}
// 		break
// 	}
// 	if len(tail) == 0 && newInterval[1] > intervals[iLen-1][1] {
// 		head[len(head)-1][1] = newInterval[1]
// 	}
// 	ret = append(head, tail...)
// 	return ret
// }
