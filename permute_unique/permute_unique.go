package permute_unique

// import "strconv"
// import "fmt"
import "sort"

func permuteUnique(nums []int) [][]int {
	// fmt.Printf("append(nums[:0], nums[1:]) = %v.\n", append(nums[:0], nums[1:]...))
	rets := [][]int{}
	if len(nums) > 0 {
		sort.Ints(nums)
		recurity(&rets, []int{}, nums)
	}
	return rets
}

func recurity(rets *[][]int, ret []int, nums []int) {
	if len(nums) == 0 {
		*rets = append(*rets, ret)
		return
	}
	for i, num := range nums {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		newNums := []int{}
		if i == 0 {
			newNums = append(newNums, nums[1:]...)
		} else if i == len(nums) - 1 {
			newNums = append(newNums, nums[:i]...)
		} else {
			newNums = append(newNums, nums[:i]...)
			newNums = append(newNums, nums[i + 1:]...)
		}
		newRet := make([]int, len(ret))
		copy(newRet, ret)
		recurity(rets, append(newRet, num), newNums)
	}
}

// func permuteUnique(nums []int) [][]int {
// 	// fmt.Printf("append(nums[:0], nums[1:]) = %v.\n", append(nums[:0], nums[1:]...))
// 	rets := [][]int{}
// 	if len(nums) > 0 {
// 		recurity(&rets, []int{}, nums)
// 	}
// 	return rets
// }

// func recurity(rets *[][]int, ret []int, nums []int) {
// 	if len(nums) == 0 {
// 		*rets = append(*rets, ret)
// 		return
// 	}
// 	rpt := make(map[int]bool)
// 	for i, num := range nums {
// 		if _, ok := rpt[num]; ok {
// 			continue
// 		} else {
// 			rpt[num] = true
// 		}
// 		newNums := []int{}
// 		if i == 0 {
// 			newNums = append(newNums, nums[1:]...)
// 		} else if i == len(nums) - 1 {
// 			newNums = append(newNums, nums[:i]...)
// 		} else {
// 			newNums = append(newNums, nums[:i]...)
// 			newNums = append(newNums, nums[i + 1:]...)
// 		}
// 		newRet := make([]int, len(ret))
// 		copy(newRet, ret)
// 		recurity(rets, append(newRet, num), newNums)
// 	}
// }

// func permuteUnique(nums []int) [][]int {
// 	// fmt.Printf("append(nums[:0], nums[1:]) = %v.\n", append(nums[:0], nums[1:]...))
// 	rets := [][]int{}
// 	if len(nums) > 0 {
// 		rpt := make(map[string]bool)
// 		recurity(&rets, []int{}, nums, &rpt, "")
// 	}
// 	return rets
// }

// func recurity(rets *[][]int, ret []int, nums []int, rpt *map[string]bool, key string) {
// 	if len(nums) == 0 {
// 		// fmt.Printf("rpt = %v, key = %s.\n", *rpt, key)
// 		if _, ok := (*rpt)[key]; !ok {
// 			*rets = append(*rets, ret)
// 			(*rpt)[key] = true
// 		}
// 		return
// 	}
// 	for i, num := range nums {
// 		newNums := []int{}
// 		if i == 0 {
// 			newNums = append(newNums, nums[1:]...)
// 		} else if i == len(nums) - 1 {
// 			newNums = append(newNums, nums[:i]...)
// 		} else {
// 			newNums = append(newNums, nums[:i]...)
// 			newNums = append(newNums, nums[i + 1:]...)
// 		}
// 		newRet := make([]int, len(ret))
// 		copy(newRet, ret)
// 		recurity(rets, append(newRet, num), newNums, rpt, key + "_" + strconv.Itoa(num))
// 	}
// }