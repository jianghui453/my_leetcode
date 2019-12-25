// 给定一个经过编码的字符串，返回它解码后的字符串。

// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

// 示例:

// s = "3[a]2[bc]", 返回 "aaabcbc".
// s = "3[a2[c]]", 返回 "accaccacc".
// s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".

package string

import (
// "fmt"
)

func decodeString(s string) string {
	l := len(s)
	ret := make([]byte, 0)

	for i := 0; i < l; i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			ret = append(ret, s[i])
		} else if s[i] >= '0' && s[i] <= '9' {
			start := i
			for ; i < l; i++ {
				if s[i] == '[' {
					i++
					break
				}
			}
			mbrackets := 1
			for ; i < l; i++ {
				if s[i] == '[' {
					mbrackets++
				} else if s[i] == ']' {
					mbrackets--
					if mbrackets == 0 {
						if i == l-1 {
							ret = append(ret, _decodeString(s[start:])...)
						} else {
							ret = append(ret, _decodeString(s[start:i+1])...)
						}
						break
					}
				}
			}
		}
	}
	return string(ret)
}

func _decodeString(s string) []byte {
	l := len(s)
	bytes := make([]byte, 0)
	cnt := 0
	i := 0
	for ; i < l && s[i] >= '0' && s[i] <= '9'; i++ {
		cnt = cnt*10 + int(s[i]-'0')
	}

	for ; i < l; i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			bytes = append(bytes, s[i])
		} else if s[i] >= '0' && s[i] <= '9' {
			start := i
			for i = i + 1; i < l && s[i] != '['; i++ {
				continue
			}

			mbrackets := 1
			for i = i + 1; i < l; i++ {
				if s[i] == '[' {
					mbrackets++
				} else if s[i] == ']' {
					mbrackets--
					if mbrackets == 0 {
						if i == l-1 {
							bytes = append(bytes, _decodeString(s[start:])...)
						} else {
							bytes = append(bytes, _decodeString(s[start:i+1])...)
						}
						break
					}
				}
			}
		}
	}

	ret := make([]byte, 0)
	for i := 0; i < cnt; i++ {
		ret = append(ret, bytes...)
	}

	return ret
}
