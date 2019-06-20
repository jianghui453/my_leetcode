package leet_code

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	var sb strings.Builder
	// var ch uint8
	var i, ch int
Loop:
	for len(strs) > 0 {
		ch = -1
		for _, s := range strs {
			if i >= len(s) {
				break Loop
			}
			if ch == -1 {
				ch = int(s[i])
				continue
			}
			if ch != int(s[i]) {
				break Loop
			}
		}
		sb.WriteByte(byte(ch))
		i ++
	}
	return sb.String()
} 