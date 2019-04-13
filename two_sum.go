package leet_code

func twoSum(nums []int, target int) []int {
	var res []int
Loop:
	for k, num := range nums {
		for ka, numa := range nums[k+1:] {
			if num+numa == target {
				res = []int{k, k + ka + 1}
				break Loop
			}
		}
	}
	return res
}
