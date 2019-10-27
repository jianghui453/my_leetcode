// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 注意空字符串可被认为是有效字符串。
// 示例 1:
// 输入: "()"
// 输出: true
// 示例 2:
// 输入: "()[]{}"
// 输出: true
// 示例 3:
// 输入: "(]"
// 输出: false
// 示例 4:
// 输入: "([)]"
// 输出: false
// 示例 5:
// 输入: "{[]}"
// 输出: true

package stack

// import (
// 	"fmt"
// )

func isValid(s string) bool {
	st := make([]uint8, 0)
	for i := range s {
		switch s[i] {
		case ')':
			if len(st) > 0 && st[len(st)-1] == '(' {
				st = st[: len(st)-1]
			} else {
				return false
			}
			break
		case '}':
			if len(st) > 0 && st[len(st)-1] == '{' {
				st = st[: len(st)-1]
			} else {
				return false
			}
			break
		case ']':
			if len(st) > 0 && st[len(st)-1] == '[' {
				st = st[: len(st)-1]
			} else {
				return false
			}
			break
		default:
			st = append(st, s[i])
		}
	}
	return len(st) == 0
}

// func isValid(s string) bool {
// 	strs := []rune{}
// 	for _, ch := range s {
// 		fmt.Printf("strs = %v; ch = %d\n", strs, ch)
// 		if ch == '}' {
// 			if len(strs) == 0 || strs[len(strs)-1] != '{' {
// 				return false
// 			}
// 			strs = strs[:len(strs)-1]
// 			continue
// 		}
// 		if ch == ']' {
// 			if len(strs) == 0 || strs[len(strs)-1] != '[' {
// 				return false
// 			}
// 			strs = strs[:len(strs)-1]
// 			continue
// 		}
// 		if ch == ')' {
// 			if len(strs) == 0 || strs[len(strs)-1] != '(' {
// 				return false
// 			}
// 			strs = strs[:len(strs)-1]
// 			continue
// 		}
// 		if ch == '{' ||
// 			ch == '[' ||
// 			ch == '(' {
// 			strs = append(strs, ch)
// 		}
// 	}
// 	if len(strs) > 0 {
// 		return false
// 	}
// 	return true
// }
