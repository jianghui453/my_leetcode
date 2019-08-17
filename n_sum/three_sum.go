//给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。
//
//例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
// [-1, 0, 1],
// [-1, -1, 2]
//]

package n_sum

func threeSum(nums []int) [][]int {
	var ret [][]int
	nLen := len(nums)
	if nLen < 3 {
		return ret
	}
	for i := 0; i < nLen - 2; i ++ {
		if i + 2 < nLen && nums[i+2] == nums[i] {

		}
	}
}

//func threeSum(nums []int) [][]int {
//	var ret [][]int
//	var rptMap = make(map[string]bool)
//
//	sort.Ints(nums)
//
//	for i := 1; i+1 < len(nums); i++ {
//		left := i - 1
//		right := i + 1
//		for left >= 0 && right < len(nums) {
//			if nums[left]+nums[i]+nums[right] > 0 {
//				left--
//				continue
//			}
//			if nums[left]+nums[i]+nums[right] < 0 {
//				right++
//				continue
//			}
//			key := fmt.Sprintf("%d-%d", nums[left], nums[right])
//			fmt.Printf("rptMap = %v; key = %s\n", rptMap, key)
//			if _, ok := rptMap[key]; !ok {
//				ret_item := []int{
//					nums[left],
//					nums[i],
//					nums[right],
//				}
//				ret = append(ret, ret_item)
//				rptMap[key] = true
//			}
//			left--
//			right++
//		}
//	}
//	return ret
//}
