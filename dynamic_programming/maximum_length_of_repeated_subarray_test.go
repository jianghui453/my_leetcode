package dynamic_programming

import (
	"testing"
)

func TestFindLength(t *testing.T) {
	var (
		A, B []int
		hope, ret int
	)

	A = []int{1,2,3,2,1}
	B = []int{3,2,1,4,7}
	hope = 3
	ret = findLength(A, B)
	t.Log(ret == hope, A, B, hope, ret)

	A = []int{1}
	B = []int{3,2,1,4,7}
	hope = 1
	ret = findLength(A, B)
	t.Log(ret == hope, A, B, hope, ret)

	A = []int{1}
	B = []int{3}
	hope = 0
	ret = findLength(A, B)
	t.Log(ret == hope, A, B, hope, ret)
}
