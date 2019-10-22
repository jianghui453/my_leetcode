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
    sl := len(s)
    if sl < 2 {
        return true
    }
    left := 0
    right := sl-1
    s = strings.ToLower(s)
    for left <= right {
        if (s[left] < 'a' || s[left] > 'z') && (s[left] < '0' || s[left] > '9') {
            left ++
            continue
        }
        if (s[right] < 'a' || s[right] > 'z') && (s[right] < '0' || s[right] > '9') {
            right --
            continue
        }
        if s[left] == s[right] {
            left ++
            right --
            continue
        }
        return false
    }
    return true
}
