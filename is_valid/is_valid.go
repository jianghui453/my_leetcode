package leet_code

import (
	"fmt"
)

func isValid(s string) bool {
	strs := []rune{}
	for _, ch := range s {
		fmt.Printf("strs = %v; ch = %d\n", strs, ch)
		if ch == '}' {
			if len(strs) == 0 || strs[len(strs)-1] != '{' {
				return false
			}
			strs = strs[:len(strs)-1]
			continue
		}
		if ch == ']' {
			if len(strs) == 0 || strs[len(strs)-1] != '[' {
				return false
			}
			strs = strs[:len(strs)-1]
			continue
		}
		if ch == ')' {
			if len(strs) == 0 || strs[len(strs)-1] != '(' {
				return false
			}
			strs = strs[:len(strs)-1]
			continue
		}
		if ch == '{' ||
			ch == '[' ||
			ch == '(' {
			strs = append(strs, ch)
		}
	}
	if len(strs) > 0 {
		return false
	}
	return true
}
