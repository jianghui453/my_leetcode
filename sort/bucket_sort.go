package sort

func bucketSort(nums []float64) {
    nLen := len(nums)
    if nLen < 2 {
        return
    }
    cntBucket := 10
    bucket := make([][]float64, cntBucket)
    for i := 0; i < nLen; i ++ {
        idxBucket := int(nums[i]*10)
        bucket[idxBucket] = append(bucket[idxBucket], nums[i])
    }
    for i := 0; i < cntBucket; i ++ {
        insertSortForBucket(bucket[i])
    }
    idx := 0
    for i := 0; i < cntBucket; i ++ {
        cnt := len(bucket[i])
        for j := 0; j < cnt; j ++ {
            nums[idx] = bucket[i][j]
            idx++
        }
    }
}

func insertSortForBucket(nums[]float64) {
    nLen := len(nums)
    for i := 1; i < nLen; i ++ {
        set := false
        for j := i-1; j >= 0; j -- {
            if nums[j] > nums[i] {
                nums[j+1] = nums[j]
            } else {
                nums[j+1] = nums[i]
                set = true
                break
            }
        }
        if !set {
            nums[0] = nums[i]
        }
    }
}