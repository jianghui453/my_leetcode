package sort

import "fmt"

func mergeSort(nums []int) {
	nLen := len(nums)
	for i := 1; i < nLen; i <<= 1 {
		for j := 0; j < nLen; j += i << 1 {
			if j+i > nLen {
				continue
			}
			nums1 := nums[j : j+i]
			var nums2 []int
			if j+i<<1 >= nLen {
				nums2 = nums[j+i:]
			} else {
				nums2 = nums[j+i : j+i<<1]
			}
			_len := len(nums1) + len(nums2)
			var _nums []int
			off1 := 0
			off2 := 0
			for off1 < i && off2 < _len-i {
				if nums1[off1] < nums2[off2] {
					_nums = append(_nums, nums1[off1])
					off1++
				} else {
					_nums = append(_nums, nums2[off2])
					off2++
				}
			}
			if off1 < i {
				_nums = append(_nums, nums1[off1:]...)
			}
			if off2 < _len-i {
				_nums = append(_nums, nums2[off2:]...)
			}

			for x := j; x < j+_len; x++ {
				nums[x] = _nums[x-j]
			}
		}
	}
}

//func mergeSort(nums []int) {
//    fmt.Printf("nums=%v\n", nums)
//    nLen := len(nums)
//    if nLen < 2 {
//        return
//    }
//    mid := nLen/2
//    nums1 := nums[:mid]
//    len1 := len(nums1)
//    nums2 := nums[mid:]
//    mergeSort(nums1)
//    mergeSort(nums2)
//    var _nums []int
//   left := 0
//   right := len1
//  for left < len1 && right < nLen {
//      if nums[left] < nums[right] {
//          _nums = append(_nums, nums[left])
//          left ++
//      } else {
//          _nums = append(_nums, nums[right])
//          right ++
//      }
//  }
//  if left < len1 {
//      _nums = append(_nums, nums[left:]...)
//  }
//  if right < nLen {
//      _nums = append(_nums, nums[right:]...)
//  }
//  for i := 0; i < nLen; i ++ {
//      nums[i] = _nums[i]
//  }
//}
