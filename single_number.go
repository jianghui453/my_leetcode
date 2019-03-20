package leet_code

func singleNumberMine(nums []int) int {
	hisNumsMap := make(map[int]bool)
	for _, num := range nums {
		if _, ok := hisNumsMap[num]; ok {
			delete(hisNumsMap, num)
			continue
		}
		hisNumsMap[num] = true
	}
	for num := range hisNumsMap {
		return num
	}
	return 0
}

func singleNumberBest(nums []int) int {
	res := 0
	for _, i := range nums {
		res ^= i
	}
	return res
}
