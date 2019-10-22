package solve_n_queens

func totalNQueens(n int) int {
	ret := 0
	if n > 0 {
		column := make([]int, n)
		strs := make([][]int, n)
		for i := 0; i < n; i++ {
			str := make([]int, n)
			strs[i] = str
		}
		totalNQueens_(n, 0, column, strs, &ret)
	}
	return ret
}

func totalNQueens_(n, y int, column []int, strs [][]int, ret *int) {
	if y >= n {
		*ret++
		return
	}
Loop:
	for x := 0; x < n; x++ {
		if column[x] > 0 {
			continue
		}
		for i := 1; x-i >= 0 && y-i >= 0; i++ {
			if strs[y-i][x-i] == 1 {
				continue Loop
			}
		}
		for i := 1; x+i < n && y-i >= 0; i++ {
			if strs[y-i][x+i] == 1 {
				continue Loop
			}
		}
		column[x]++
		strs[y][x]++
		totalNQueens_(n, y+1, column, strs, ret)
		column[x]--
		strs[y][x]--
	}
}
