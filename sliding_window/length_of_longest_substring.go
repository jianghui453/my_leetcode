//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
//示例 1:
//
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

//标签：哈希表、双指针、字符串、滑动窗口

package sliding_window

func lengthOfLongestSubstring(s string) int {
    sLen := len(s)
    if sLen < 2 {
        return sLen
    }
    his := make([]int, 255)
    left := 0
    right := 1
    his[s[0]] = 1
    max := 1
    for right < sLen {
        if his[s[right]] == 0 {
            if max < (right-left+1) {
                max = right-left+1
            }
        } else {
            for left < his[s[right]] {
                his[s[left]] = 0
                left ++
            }
        }
        his[s[right]] = right+1
        right++
    }
    return max
}

//func lengthOfLongestSubstring(s string) int {
//	var max int
//	smap := make(map[rune]int)
//	count := 0
//	for offset, ch := range s {
//		i := -1
//		if _, ok := smap[ch]; ok {
//			i = smap[ch]
//		}
//		if i >= offset-count {
//			smap[ch] = offset
//			count = offset - i
//			continue
//		}
//		smap[ch] = offset
//		count++
//		if count > max {
//			max = count
//		}
//	}
//	return max
//}

//func lengthOfLongestSubstring(s string) int {
//	var max, count, offset int
//	smap := make(map[uint8]int)
//	slen := len(s)
//	for slen > 0 {
//		if offset >= slen {
//			break
//		}
//		if i, ok := smap[s[offset]]; ok {
//			for j := offset - count; j <= i; j++ {
//				delete(smap, s[j])
//			}
//			count = offset - i
//			smap[s[offset]] = offset
//			offset++
//			continue
//		}
//		smap[s[offset]] = offset
//		count++
//		offset++
//		if max < count {
//			max = count
//		}
//	}
//	return max
//}
