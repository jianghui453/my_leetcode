package sort

func shellSort(nums []int) {
	nLen := len(nums)
	for h := nLen / 2; h > 0; h /= 2 {
		for i := h; i < nLen; i++ {
			insertI(nums, h, i)
		}
	}
}

func insertI(nums []int, h, i int) {
	num := nums[i]
	exchange := false
	for j := i - h; j >= 0; j -= h {
		if nums[j] > num {
			nums[j+h] = nums[j]
		} else {
			nums[j+h] = num
			exchange = true
			break
		}
	}
	if !exchange {
		nums[i%h] = num
	}
}
