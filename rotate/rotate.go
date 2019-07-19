package rotate

func rotate(matrix [][]int) {
	var ret = [][]int{}
	mLen := len(matrix)
	for i := 0; i < mLen; i ++ {
		nums := []int{}
		for j := mLen - 1; j >= 0; j -- {
			nums = append(nums, matrix[j][i])
		}
		ret = append(ret, nums)
	}
	copy(matrix, ret)
}