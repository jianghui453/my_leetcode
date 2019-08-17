package merge

func insert(intervals [][]int, newInterval []int) [][]int {
	var ret [][]int
	iLen := len(intervals)
	if iLen == 0 {
		return append(ret, newInterval)
	}
	if newInterval[1] < intervals[0][0] {
		ret = append(ret, newInterval)
		ret = append(ret, intervals...)
		return ret
	}
	if newInterval[0] > intervals[iLen-1][1] {
		ret = append(ret, intervals...)
		ret = append(ret, newInterval)
		return ret
	}
	var head, tail [][]int
	for i := 0; i < iLen; i++ {
		if newInterval[0] > intervals[i][0] {
		    head = append(head, intervals[i])
			continue
		}
        if i == 0 {
            head = append(head, newInterval)
            break
        }
		if newInterval[0] > intervals[i-1][1] {
		    head = append(head, newInterval)
		    break
        }
		if newInterval[1] > intervals[i-1][1] {
		    head[len(head)-1][1] = newInterval[1]
        }
	}
	for i := iLen-1; i >= 0; i -- {
	    if intervals[i][0] > newInterval[1] {
	        tail = append([][]int{{intervals[i][0], intervals[i][1]}}, tail...)
	        continue
        }
	    if intervals[i][1] > newInterval[1] {
	        head[len(head)-1][1] = intervals[i][1]
        }
	    break
    }
	if len(tail) == 0 && newInterval[1] > intervals[iLen-1][1]  {
	    head[len(head)-1][1] = newInterval[1]
    }
	ret = append(head, tail...)
	return ret
}
