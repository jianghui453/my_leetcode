// 根据每日 气温 列表，请重新生成一个列表，对应位置的输入是你需要再等待多久温度才会升高超过该日的天数。如果之后都不会升高，请在该位置用 0 来代替。

// 例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。

// 提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。

package stack

import (
	"fmt"
)

func dailyTemperatures(T []int) []int {
	var (
		l int = len(T)
		ret []int = make([]int, l)
		s []int = make([]int, 0)
	)

	for i := 0; i < l; i++ {
		fmt.Println(s)
		if len(s) == 0 {
			ret[i] = 0
			s = append(s, i)
		} else {
			sl := len(s)-1
			for sl >= 0 && T[i] > T[s[sl]] {
				ret[s[sl]] = i-s[sl]
				sl--
			}

			if sl == len(s)-1 {
				s = append(s, i)
			} else if sl >= 0 {
				s = append(s[: sl+1], i)
			} else {
				s = []int{i}
			}
		}
	}

	return ret
}