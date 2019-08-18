package unique_path

func uniquePaths(m int, n int) int {
	if m < 2 || n < 2 {
		return 1
	}
	dp := make([][]int, m)
	for i := 0; i < m; i ++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for _m := 0; _m < m; _m ++ {
		for _n := 0; _n < n; _n ++ {
			if _n > 0 && _m > 0 {
				dp[_m][_n] = dp[_m-1][_n] + dp[_m][_n-1]
			} else if _n == 0 && _m > 0 {
				dp[_m][_n] = dp[_m-1][_n]
			} else if _n > 0 && _m == 0 {
				dp[_m][_n] = dp[_m][_n-1]
			}
		}
	}
	return dp[m-1][n-1]
}

//func uniquePaths(m int, n int) int {
//	if m < 2 && n < 2 {
//		return 0
//	}
//	ret := 0
//	_uniquePaths(m, n, 1, 1, &ret)
//	return ret
//}
//
//func _uniquePaths(m, n, _m, _n int, ret *int) {
//	if _m >= m && _n >= n {
//		*ret++
//		return
//	}
//	if _m < m {
//		_uniquePaths(m, n, _m+1, _n, ret)
//	}
//	if _n < n {
//		_uniquePaths(m, n, _m, _n+1, ret)
//	}
//}