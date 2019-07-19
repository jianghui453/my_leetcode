package combination_sum

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var ret [][]int
	var his []int
	sum(&ret, his, candidates, target)
	return ret	
}

func sum(ret *[][]int, his []int, candidates []int, target int) {
	for i := 0; i < len(candidates); i ++ {
		if i > 0 && candidates[i] == candidates[i - 1] {
			continue
		}
		if candidates[i] < target && i < len(candidates) - 1 {
			newHis := make([]int, len(his) + 1)
			copy(newHis, his)
			newHis[len(his)] = candidates[i]
			sum(ret, newHis, candidates[i + 1:], target - candidates[i])
		} else if candidates[i] == target {
			newHis := make([]int, len(his) + 1)
			copy(newHis, his)
			newHis[len(his)] = candidates[i]
			(*ret) = append((*ret), newHis)
			break
		} else {
			break
		}
	}
}