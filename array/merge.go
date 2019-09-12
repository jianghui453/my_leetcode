//给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
//
//说明:
//
//初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
//你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
//示例:
//
//输入:
//nums1 = [1,2,3,0,0,0], m = 3
//nums2 = [2,5,6],       n = 3
//
//输出: [1,2,2,3,5,6]

package array

func merge(nums1 []int, m int, nums2 []int, n int) {
	i1 := 0
	i2 := 0
	nums := make([]int, m+n, m+n)
	i0 := 0
	for i1 < m && i2 < n {
		if nums1[i1] > nums2[i2] {
			nums[i0] = nums2[i2]
			i2++
		} else {
			nums[i0] = nums1[i1]
			i1++
		}
		i0++
	}
	if i1 < m {
		for ; i1 < m; i1++ {
			nums[i0] = nums1[i1]
			i0++
		}
	}
	if i2 < n {
		for ; i2 < n; i2++ {
			nums[i0] = nums2[i2]
			i0++
		}
	}
	copy(nums1, nums)
}
