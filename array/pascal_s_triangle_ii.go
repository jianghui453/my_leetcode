//给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
//
//
//
//在杨辉三角中，每个数是它左上方和右上方的数的和。
//
//示例:
//
//输入: 3
//输出: [1,3,3,1]
//进阶：
//
//你可以优化你的算法到 O(k) 空间复杂度吗？

package array

func getRow(rowIndex int) []int {
	ret := []int{1}
	for ; rowIndex > 0; rowIndex-- {
		newRet := make([]int, 0)
		newRet = append(newRet, ret[0])
		l := len(ret)
		for i := 0; i < l-1; i++ {
			newRet = append(newRet, ret[i]+ret[i+1])
		}
		newRet = append(newRet, ret[l-1])
		ret = newRet
	}

	return ret
}

// func getRow(rowIndex int) []int {
// 	r := make([]int, 0)
// 	r = []int{1}
// 	if rowIndex == 0 {
// 		return r
// 	}
// 	r = []int{1, 1}
// 	if rowIndex == 1 {
// 		return r
// 	}
// 	lenR := 2
// 	for i := 2; i <= rowIndex; i++ {
// 		s := []int{1}
// 		lenS := 1
// 		for j := 0; j < lenR-1; j++ {
// 			s = append(s, r[j]+r[j+1])
// 			lenS++
// 		}
// 		s = append(s, 1)
// 		lenS++
// 		r = s
// 		lenR = lenS
// 	}
// 	return r
// }