package backtracking

import "testing"

func TestGrayCode(t *testing.T) {
	var n int
	var hope, ret []int

	//n = 2
	//hope = []int{0, 1, 3, 2}
	//ret = grayCode(n)
	//t.Logf("\nhope=%v \n ret=%v", hope, ret)

	n = 3
	hope = []int{0, 1, 3, 2, 6, 7, 5, 4}
	ret = grayCode(n)
	t.Logf("\nhope=%v \n ret=%v", hope, ret)
}
