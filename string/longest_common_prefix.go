// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。
// 示例 1:
// 输入: ["flower","flow","flight"]
// 输出: "fl"
// 示例 2:
// 输入: ["dog","racecar","car"]
// 输出: ""
// 解释: 输入不存在公共前缀。
// 说明:
// 所有输入只包含小写字母 a-z 。

package string

// import "fmt"

func longestCommonPrefix(strs []string) string {
	strsLen := len(strs)
	if strsLen == 0 {
		return ""
	} else if strsLen == 1 {
		return strs[0]
	}
	prefix := ""
loop:
	for i := range strs[0] {
		for j := 1; j < strsLen; j++ {
			if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
				break loop
			}
		}
		prefix += string(strs[0][i])
	}
	return prefix
}

// func longestCommonPrefix(strs []string) string {
// 	var sb strings.Builder
// 	// var ch uint8
// 	var i, ch int
// Loop:
// 	for len(strs) > 0 {
// 		ch = -1
// 		for _, s := range strs {
// 			if i >= len(s) {
// 				break Loop
// 			}
// 			if ch == -1 {
// 				ch = int(s[i])
// 				continue
// 			}
// 			if ch != int(s[i]) {
// 				break Loop
// 			}
// 		}
// 		sb.WriteByte(byte(ch))
// 		i ++
// 	}
// 	return sb.String()
// }
