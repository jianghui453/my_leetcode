package array

//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
//
//说明：
//
//你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
//
//示例 1:
//
//输入: [2,2,1]
//输出: 1
//示例 2:
//
//输入: [4,1,2,1,2]
//输出: 4

//func singleNumber(nums []int) int {
//	hisNumsMap := make(map[int]bool)
//	for _, num := range nums {
//		if _, ok := hisNumsMap[num]; ok {
//			delete(hisNumsMap, num)
//			continue
//		}
//		hisNumsMap[num] = true
//	}
//	for num := range hisNumsMap {
//		return num
//	}
//	return 0
//}

//func singleNumber(nums []int) int {
//	res := 0
//	for _, i := range nums {
//		res ^= i
//	}
//	return res
//}

//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。
//
//说明：
//
//你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
//
//示例 1:
//
//输入: [2,2,3,2]
//输出: 3
//示例 2:
//
//输入: [0,1,0,1,0,1,99]
//输出: 99

func singleNumber(nums []int) int {
	lenNums := len(nums)

	// a b  nums[i]  a b
	// 0 0  0        0 0
	// 1 0  1        0 0

	// 1 0  0        1 0
	// 0 1  1        1 0

	// 0 1  0        0 1
	// 0 0  1        0 1
	var a, b int
	for i := 0; i < lenNums; i ++ {
		c := nums[i]
		_a := (a^b)&(a^c)
		b = (a^-1)&(b^c)
		a = _a
	}
	return a|b
}
