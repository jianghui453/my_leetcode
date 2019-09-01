//给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
//
//说明：解集不能包含重复的子集。
//
//示例:
//
//输入: nums = [1,2,3]
//输出:
//[
//  [3],
//  [1],
//  [2],
//  [1,2,3],
//  [1,3],
//  [2,3],
//  [1,2],
//  []
//]

package array

func subSets(nums []int) [][]int {
	lenN := len(nums)
	ret := [][]int{{}}
	if lenN < 1 {
		return ret
	}
	for i := 0; i < lenN; i++ {
		getSubsets(nums[i:], []int{}, &ret)
	}
	return ret
}

func getSubsets(nums, numsR []int, ret *[][]int) {
	lenN := len(nums)
	if lenN == 0 {
		return
	}
	numsR = append(numsR, nums[0])
	lenNR := len(numsR)
	*ret = append(*ret, numsR)
	for i := 1; i < lenN; i++ {
		numsRNew := make([]int, lenNR)
		copy(numsRNew, numsR)
		getSubsets(nums[i:], numsRNew, ret)
	}
}
