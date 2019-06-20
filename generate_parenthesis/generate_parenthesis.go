package leet_code

// import(
// 	"fmt"
// 	"strings"
// )

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	var strs, strs_new []string
	var left, right int
	strs = []string{""}
Loop:
	for {
		strs_new = []string{}
		for _, str := range strs {
			left = 0
			right = 0
			for _, ch := range str {
				if ch == '(' {
					left++
				} else if ch == ')' {
					right++
				}
			}
			if left == n && right == n {
				break Loop
			}
			if left < n {
				strs_new = append(strs_new, str+"(")
			}
			if right < left {
				strs_new = append(strs_new, str+")")
			}
		}
		strs = strs_new
	}
	return strs
}
