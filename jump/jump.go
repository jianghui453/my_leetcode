package jump

import "fmt"

func jump(nums []int) int {
	nLen := len(nums)
	if nLen < 2 {
		return 0
	}
	idx := 0
	step := 0
	for idx < nLen {
		fmt.Printf("idx = %d.\n", idx)
		if idx + nums[idx] >= nLen - 1 {
			return step + 1
		}
		max := idx
		for i := idx + 1; i <= idx + nums[idx]; i ++ {
			fmt.Printf("i = %d, nums[i] = %d.\n", i, nums[i])
			if i + nums[i] > max + nums[max] {
				max = i
			}
		}
		idx = max
		step ++
	}
	return step
}