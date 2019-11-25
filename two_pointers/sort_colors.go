//给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
//
//此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
//
//注意:
//不能使用代码库中的排序函数来解决这道题。
//
//示例:
//
//输入: [2,0,2,1,1,0]
//输出: [0,0,1,1,2,2]
//进阶：
//
//一个直观的解决方案是使用计数排序的两趟扫描算法。
//首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。
//你能想出一个仅使用常数空间的一趟扫描算法吗？

package two_pointers

import (
// "fmt"
)

func sortColors(nums []int) {
	l := len(nums)
	if l <= 1 {
		return
	}

	i, i0, i1, i2 := 0, 0, 0, l-1

	for i <= i2 {
		switch nums[i] {
		case 0:
			if i0 == i1 {
				nums[i0], nums[i] = nums[i], nums[i0]
				i0++
				i1++
			} else {
				nums[i1], nums[i] = nums[i], nums[i1]
				nums[i0], nums[i1] = nums[i1], nums[i0]
				i0++
				i1++
			}
		case 1:
			nums[i0], nums[i] = nums[i], nums[i0]
			i1++
		case 2:
			nums[i2], nums[i] = nums[i], nums[i2]
			i2--
		}

		if i < i1 {
			i++
		}
	}

	return
}

// func sortColors(nums []int) {
// 	nLen := len(nums)
// 	if nLen < 1 {
// 		return
// 	}
// 	cur := 0
// 	zero := 0
// 	two := nLen - 1
// 	for cur <= two {
// 		fmt.Printf("nums=%v cur=%d zero=%d two=%d\n", nums, cur, zero, two)
// 		switch nums[cur] {
// 		case 0:
// 			nums[cur], nums[zero] = nums[zero], nums[cur]
// 			zero++
// 			if zero >= cur {
// 				cur++
// 			}
// 			break
// 		case 1:
// 			cur++
// 			break
// 		case 2:
// 			nums[cur], nums[two] = nums[two], nums[cur]
// 			two--
// 			break
// 		}
// 	}
// }

//func sortColors(nums []int) {
//    nLen := len(nums)
//    left := 0
//    right := nLen - 1
//    for i := 0; i < nLen; i ++ {
//        switch nums[i] {
//        case 0:
//            left++
//            break
//        case 2:
//            right--
//            break
//        default:
//            break
//        }
//    }
//    for i := 0; i < left; i ++ {
//        nums[i] = 0
//    }
//    for i := left; i <= right; i ++ {
//        nums[i] = 1
//    }
//    for i := right + 1; i < nLen; i ++ {
//        nums[i] = 2
//    }
//}
