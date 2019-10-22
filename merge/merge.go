package merge

func merge(intervals [][]int) [][]int {
	iLen := len(intervals)
	if iLen > 1 {
		for {
			intervals = _merge(intervals)
			_iLen := len(intervals)
			if _iLen == iLen {
				break
			}
			iLen = _iLen
		}

	}
	return intervals
}

func _merge(intervals [][]int) [][]int {
	var ret [][]int
	iLen := len(intervals)
Loop:
	for i := 0; i < iLen; i++ {
		for j := 0; j < len(ret); j++ {
			if ret[j][1] < intervals[i][0] || ret[j][0] > intervals[i][1] {
				continue
			}
			var min, max int
			if ret[j][0] < intervals[i][0] {
				min = ret[j][0]
			} else {
				min = intervals[i][0]
			}
			if ret[j][1] > intervals[i][1] {
				max = ret[j][1]
			} else {
				max = intervals[i][1]
			}
			ret[j] = []int{min, max}
			continue Loop
		}
		ret = append(ret, intervals[i])
	}
	return ret
}
