package linked_list

import (
	"testing"
)

func Test_SkipListInsert(t *testing.T) {
	var (
		sl *SkipList
	)

	sl = NewSkipList(3, 0.5)

	sl.Insert(3)
	t.Log(sl, toSlice(sl))

	sl.Insert(6)
	t.Log(sl, toSlice(sl))

	sl.Insert(7)
	t.Log(sl, toSlice(sl))

	sl.Insert(9)
	t.Log(sl, toSlice(sl))

	sl.Insert(12)
	t.Log(sl, toSlice(sl))

	sl.Insert(19)
	t.Log(sl, toSlice(sl))

	sl.Insert(17)
	t.Log(sl, toSlice(sl))

	sl.Insert(26)
	t.Log(sl, toSlice(sl))

	sl.Insert(21)
	t.Log(sl, toSlice(sl))

	sl.Insert(25)
	t.Log(sl, toSlice(sl))

	t.Log(sl, sl.Search(9), toSlice(sl))

	sl.Delete(9)
	t.Log(sl, toSlice(sl))

	sl.Delete(3)
	t.Log(sl, toSlice(sl))

	sl.Delete(6)
	t.Log(sl, toSlice(sl))

	sl.Delete(26)
	t.Log(sl, toSlice(sl))
}

func toSlice(sl *SkipList) [][]int {
	if sl.Header == nil {
		return [][]int{}
	}

	ret := make([][]int, sl.Level+1)
	for i := 0; i <= sl.Level; i++ {
		ret[i] = make([]int, 0)
		cur := sl.Header
		for ; cur != nil; cur = cur.Forword[i] {
			ret[i] = append(ret[i], cur.Key)
		}
	}

	return ret
}
