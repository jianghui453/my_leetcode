package sort

import (
    "fmt"
    "math"
)

func radioSort(nums []int) {
    nLen := len(nums)
    if nLen < 2 {
        return
    }

    cntBucket := 10
    for i := 1; i < math.MaxInt64; i *= 10 {
        bucket := make([][]int, cntBucket)
        fmt.Println()
        for j := 0; j < nLen; j ++ {
            idxBucket := nums[j]/i%10
            fmt.Printf("idxBucket=%d nums[j]=%d i=%d\n", idxBucket, nums[j], i)
            bucket[idxBucket] = append(bucket[idxBucket], nums[j])
        }
        if len(bucket[0]) == nLen {
            break
        }
        idx := 0
        for x := 0; x < cntBucket; x ++ {
            lenBucket := len(bucket[x])
            for y := 0; y < lenBucket; y ++ {
                fmt.Printf("idx=%d x=%d y=%d bucket[x]=%v\n", idx, x, y, bucket[x])
                nums[idx] = bucket[x][y]
                idx++
            }
        }
    }
}