package solve_n_queens

func solveNQueens(n int) [][]string {
    var ret [][]string
    if n > 0 {
        var b []byte
        for i := 0; i < n; i++ {
            b = append(b, '.')
        }
        var strs [][]byte
        column := make([]int, n)
        for i := 0; i < n; i ++ {
            b_ := make([]byte, n)
            copy(b_, b)
            strs = append(strs, b_)
        }
        solve(0, n, column, &ret, strs)
    }
    return ret
}

func solve(y, n int, column []int, ret *[][]string, strs [][]byte) {
    //fmt.Printf("y=%d colume=%v strs=%v\n", y, column, strs)
    if y >= n {
        ret_ := make([]string, n)
        for i, v := range strs {
            ret_[i] = string(v)
        }
        *ret = append(*ret, ret_)
        return
    }
    Loop:
    for x := 0; x < n; x++ {
        if column[x] > 0 {
            continue
        }
        i := 1
        for ; x - i >= 0 && y - i >= 0; i ++ {
            if strs[y-i][x-i] == 'Q' {
                continue Loop
            }
        }
        i = 1
        for ; x + i < n && y - i >= 0; i ++ {
            if strs[y-i][x+i] == 'Q' {
                continue Loop
            }
        }
        column[x] ++
        strs[y][x] = 'Q'
        solve(y+1, n, column, ret, strs)
        column[x] --
        strs[y][x] = '.'
    }
    //fmt.Println()
}