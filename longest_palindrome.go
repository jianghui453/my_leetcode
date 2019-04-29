package leet_code

func longestPalindrome(s string) string {
	ret := ""
	strLen := len(s)
	start := 0
	for {
		end := start
		for {
			if end < strLen && s[end] == s[start] {
				end++
				continue
			}
			end--
			break
		}
		i := 1
		for {
			if start-i >= 0 && end+i < strLen && s[start-i] == s[end+i] {
				i++
				continue
			}
			i--
			break
		}
		if end+i-(start-i)+1 > len(ret) {
			ret = s[start-i : end+i+1]
		}
		if end+1 < strLen {
			start = end + 1
			continue
		}
		break
	}
	return ret
}
