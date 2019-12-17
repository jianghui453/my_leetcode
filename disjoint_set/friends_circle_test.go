package disjoint_set

import (
	"testing"
)

func Test_findCircleNum(t *testing.T) {
	var M [][]int
	var hope, ret int

	M = [][]int{
		{1,1,0},
 		{1,1,0},
 		{0,0,1},
	}
	hope = 2
	ret = findCircleNum(M)
	t.Log(ret == hope, M, hope, ret)

	M = [][]int{
		{1,1,0},
		{1,1,1},
  		{0,1,1},
	}
	hope = 1
	ret = findCircleNum(M)
	t.Log(ret == hope, M, hope, ret)
}