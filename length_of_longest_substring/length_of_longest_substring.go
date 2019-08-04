// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

package length_of_longest_substring

func lengthOfLongestSubstring(s string) int {
	var max int
	smap := make(map[rune]int)
	count := 0
	for offset, ch := range s {
		i := -1
		if _, ok := smap[ch]; ok {
			i = smap[ch]
		}
		if i >= offset-count {
			smap[ch] = offset
			count = offset - i
			continue
		}
		smap[ch] = offset
		count++
		if count > max {
			max = count
		}
	}
	return max
}

func lengthOfLongestSubstringV1(s string) int {
	var max, count, offset int
	smap := make(map[uint8]int)
	slen := len(s)
	for slen > 0 {
		if offset >= slen {
			break
		}
		if i, ok := smap[s[offset]]; ok {
			for j := offset - count; j <= i; j++ {
				delete(smap, s[j])
			}
			count = offset - i
			smap[s[offset]] = offset
			offset++
			continue
		}
		smap[s[offset]] = offset
		count++
		offset++
		if max < count {
			max = count
		}
	}
	return max
}
