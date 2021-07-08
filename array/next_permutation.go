// 实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
//
//如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
//
//必须原地修改，只允许使用额外常数空间。
//
//以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
//1,2,3 → 1,3,2
//3,2,1 → 1,2,3
//1,1,5 → 1,5,1

package array

func nextPermutation(nums []int) {
	numsLen := len(nums)
	if numsLen <= 1 {
		return
	}

	reverseNums := func(_nums []int) {
		_numsLen := len(_nums)
		if _numsLen <= 1 {
			return
		}
		_halfNumsLen := _numsLen / 2
		for j := 0; j < _halfNumsLen; j++ {
			_nums[j], _nums[_numsLen-1-j] = _nums[_numsLen-1-j], _nums[j]
		}
	}

	for i := numsLen - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			j := i + 1
			for ; j < numsLen; j++ {
				if nums[i] >= nums[j] {
					break
				}
			}
			nums[i], nums[j-1] = nums[j-1], nums[i]
			reverseNums(nums[i+1:])
			return
		}
	}

	reverseNums(nums)
}

//func nextPermutation(nums []int) {
//	if len(nums) == 0 {
//		return
//	}
//	for i := len(nums) - 2; i >= 0; i-- {
//		if nums[i] >= nums[i+1] {
//			continue
//		}
//		fmt.Printf("i = %d\n", i)
//		var j int
//		for j = i + 1; j < len(nums) - 1; j++ {
//			if nums[j] <= nums[i] {
//				break
//			}
//		}
//		fmt.Printf("j = %d\n", j)
//		if j == len(nums) - 1 && nums[j] > nums[i] {
//			fmt.Printf("before = %v\n", nums)
//			nums[i], nums[j] = nums[j], nums[i]
//			fmt.Printf("after = %v\n", nums)
//		} else {
//			nums[i], nums[j - 1] = nums[j - 1], nums[i]
//		}
//		if i == len(nums) - 2 {
//			return
//		}
//		for j = (len(nums) - 1 - i) / 2; j > 0; j-- {
//			fmt.Printf("j = %d\v", j)
//			nums[i+j], nums[len(nums)-j] = nums[len(nums)-j], nums[i+j]
//		}
//		return
//	}
//	for j := len(nums) / 2 - 1 ; j >= 0; j -- {
//		nums[j], nums[len(nums) - 1 - j] = nums[len(nums) - 1 - j], nums[j]
//	}
//}
