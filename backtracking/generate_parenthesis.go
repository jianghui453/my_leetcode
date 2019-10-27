// 给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
// 例如，给出 n = 3，生成结果为：
// [
//   "((()))",
//   "(()())",
//   "(())()",
//   "()(())",
//   "()()()"
// ]

package backtracking

// import(
// 	"fmt"
	// "strings"
// )

func generateParenthesis(n int) []string {
	ret := make([]string, 0)

	var bt func (_n, cnt int, s string)
	bt = func(_n, cnt int, s string) {
		if _n > 0 {
			bt(_n-1, cnt+1, s+"(")
		}
		if cnt > 0 {
			if _n == 0 && cnt == 1 {
				ret = append(ret, s+")")
			} else {
				bt(_n, cnt-1, s+")")
			}
		}
	}

	bt(n, 0, "")
	return ret
}

// func generateParenthesis(n int) []string {
// 	if n == 0 {
// 		return []string{}
// 	}
// 	var strs, strs_new []string
// 	var left, right int
// 	strs = []string{""}
// Loop:
// 	for {
// 		strs_new = []string{}
// 		for _, str := range strs {
// 			left = 0
// 			right = 0
// 			for _, ch := range str {
// 				if ch == '(' {
// 					left++
// 				} else if ch == ')' {
// 					right++
// 				}
// 			}
// 			if left == n && right == n {
// 				break Loop
// 			}
// 			if left < n {
// 				strs_new = append(strs_new, str+"(")
// 			}
// 			if right < left {
// 				strs_new = append(strs_new, str+")")
// 			}
// 		}
// 		strs = strs_new
// 	}
// 	return strs
// }
