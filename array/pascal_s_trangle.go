//给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
//
//
//
//在杨辉三角中，每个数是它左上方和右上方的数的和。
//
//示例:
//
//输入: 5
//输出:
//[
//     [1],
//    [1,1],
//   [1,2,1],
//  [1,3,3,1],
// [1,4,6,4,1]
//]

package array

func generate(numRows int) [][]int {
	ret := make([][]int, 0)
	if numRows == 0 {
		return ret
	}
	ret = append(ret, []int{1})
	for ; numRows > 1; numRows-- {
		retLen := len(ret)
		retILen := len(ret[retLen-1])

		retItem := make([]int, 0)
		retItem = append(retItem, ret[retLen-1][0])
		for i := 0; i < retILen-1; i++ {
			retItem = append(retItem, ret[retLen-1][i]+ret[retLen-1][i+1])
		}
		retItem = append(retItem, ret[retLen-1][retILen-1])
		
		ret = append(ret, retItem)
	}

	return ret
}

// func generate(numRows int) [][]int {
// 	r := make([][]int, 0)
// 	if numRows == 0 {
// 		return r
// 	}
// 	r = append(r, []int{1})
// 	if numRows == 1 {
// 		return r
// 	}
// 	r = append(r, []int{1, 1})
// 	if numRows == 2 {
// 		return r
// 	}
// 	lenR := 2
// 	for i := 2; i < numRows; i++ {
// 		s := []int{1}
// 		lenS := 1
// 		for j := 0; j < lenR-1; j++ {
// 			s = append(s, r[i-1][j]+r[i-1][j+1])
// 			lenS++
// 		}
// 		s = append(s, 1)
// 		lenS++
// 		r = append(r, s)
// 		lenR = lenS
// 	}
// 	return r
// }

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
