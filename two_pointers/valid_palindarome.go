//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//
//说明：本题中，我们将空字符串定义为有效的回文串。
//
//示例 1:
//
//输入: "A man, a plan, a canal: Panama"
//输出: true
//示例 2:
//
//输入: "race a car"
//输出: false

package two_pointers

import (
	"strings"
)

func isPalindrome(s string) bool {
	l := len(s)
	if l < 2 {
		return true
	}
	left, right := 0, l-1
	s = strings.ToLower(s)
	for left <= right {
		if (s[left] < 'a' || s[left] > 'z') && (s[left] < '0' || s[left] > '9') {
			left++
		} else if (s[right] < 'a' || s[right] > 'z') && (s[right] < '0' || s[right] > '9') {
			right--
		} else if s[left] == s[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}
