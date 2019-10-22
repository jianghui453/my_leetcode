package combination_sum

import "testing"

func TestCombinationSum(t *testing.T) {
	var (
		candidates []int
		target     int
		hope       [][]int
		ret        [][]int
	)

	candidates = []int{2, 3, 6, 7}
	target = 7
	hope = [][]int{
		[]int{7},
		[]int{2, 2, 3},
	}
	ret = combinationSum(candidates, target)
	t.Logf("candidate = %v\ntarget = %d\nhope = %v\nret = %v\n\n", candidates, target, hope, ret)

	candidates = []int{7, 3, 2}
	target = 18
	ret = combinationSum(candidates, target)
	t.Logf("candidate = %v\ntarget = %d\nhope = [[2,2,2,2,2,2,2,2,2],[2,2,2,2,2,2,3,3],[2,2,2,2,3,7],[2,2,2,3,3,3,3],[2,2,7,7],[2,3,3,3,7],[3,3,3,3,3,3]]\nret = %v\n\n", candidates, target, ret)
}
