package permute

// import "fmt"

func permute(nums []int) [][]int {
	// fmt.Printf("append(nums[:0], nums[1:]) = %v.\n", append(nums[:0], nums[1:]...))
	rets := [][]int{}
	if len(nums) > 0 {
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
		newNums := []int{}
		if i == 0 {
			newNums = append(newNums, nums[1:]...)
		} else if i == len(nums)-1 {
			newNums = append(newNums, nums[:i]...)
		} else {
			newNums = append(newNums, nums[:i]...)
			newNums = append(newNums, nums[i+1:]...)
		}
		newRet := make([]int, len(ret))
		copy(newRet, ret)
		recurity(rets, append(newRet, num), newNums)
	}
}
