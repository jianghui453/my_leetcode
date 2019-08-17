package sort

func quickSort(nums []int) {
	nLen := len(nums)
	if nLen < 2 {
		return
	}
	num := nums[0]
	left := 0
	right := nLen - 1
	_nums := make([]int, nLen)
	copy(_nums, nums)
	for i := 1; i < nLen; i++ {
		if _nums[i] < num {
			nums[left] = _nums[i]
			left++
		} else {
			nums[right] = _nums[i]
			right--
		}
	}
	nums[left] = num
	//fmt.Printf("nums=%v \nnums[:left]=%v \nnums[right:]=%v\n", nums, nums[:left], nums[right:])
	quickSort(nums[:left])
	quickSort(nums[right+1:])
}
