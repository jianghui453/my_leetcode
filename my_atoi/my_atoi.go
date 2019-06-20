package leet_code

import (
	"math"
	"fmt"
)

// 第一个非空字符必须为数字或者正负号
// 正负号后面必须跟着非0数字
func myAtoi(str string) int {
	var ret, flag int
	for i := 0; i < len(str); i ++ {
		fmt.Printf("ret=%d; flag=%d; ch=%v\n", ret, flag, str[i])
		// 已经越界就退出
		if ret > math.MaxInt32 + 1 {
			break
		}
		// 负号
		if str[i] == '-' && (i == 0 || str[i - 1] == ' ') {
			flag = -1
			continue
		}
		// 正号
		if str[i] == '+' && (i == 0 || str[i - 1] == ' ') {
			flag = 1
			continue
		}
		// 空串
		if str[i] == ' ' && flag == 0 {
			continue
		}
		// 数字
		if str[i] >= '0' && str[i] <= '9' {
			ret = ret * 10 + (int(str[i]) - '0')
			if flag == 0 {
				flag = 1
			}
			continue
		}
		break
	}
	fmt.Printf("ret=%d; flag=%d\n", ret, flag)
	if flag == 1 && ret > math.MaxInt32 {
		return math.MaxInt32
	}
	if flag == -1 && 0 - ret < math.MinInt32 {
		return math.MinInt32
	}
	return ret * flag
}
