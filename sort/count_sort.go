package sort

func countSort(nums []int) {
    nLen := len(nums)
    if nLen < 2 {
        return
    }
    maxVal := nums[0]
    minVal := nums[0]
    for i := 0; i < nLen; i ++ {
        if nums[i] > maxVal {
            maxVal = nums[i]
        } else if nums[i] < minVal {
            minVal = nums[i]
        }
    }

    tLen := maxVal - minVal + 1
    tNums := make([]int, tLen)
    for i := 0; i < nLen; i ++ {
        tNums[nums[i]-minVal] ++
    }
    idx := 0
    for i := 0; i < tLen; i ++ {
        for j := 0; j < tNums[i]; j ++ {
            nums[idx] = i
            idx ++
        }
    }
}