// 给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
// 注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。
//
// 示例 1：
// 输入：
//   s = "barfoothefoobarman",
//   words = ["foo","bar"]
// 输出：[0,9]
// 解释：
// 从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
// 输出的顺序不重要, [9,0] 也是有效答案。
// 示例 2：
// 输入：
//   s = "wordgoodgoodgoodbestword",
//   words = ["word","good","best","word"]
// 输出：[]

package find_sub_string

// import (
// 	"strings"
// 	"fmt"
// )

func findSubstring(s string, words []string) []int {
	wordsLen := len(words)
	sLen := len(s)
	ret := make([]int, 0)
	if wordsLen == 0 {
		return ret
	}

	wordsMap := make(map[string]int)
	for i := 0; i < wordsLen; i++ {
		if _, ok := wordsMap[words[i]]; !ok {
			wordsMap[words[i]] = 1
		} else {
			wordsMap[words[i]]++
		}
	}

	wordLen := len(words[0])
	for i := 0; i < wordLen; i++ {

		winMap := make(map[string]int)
		cnt := 0
		left, right := i, i
		for left <= right && right <= sLen-wordLen {

			var subStr string
			if right+wordLen == sLen {
				subStr = s[right:]
			} else {
				subStr = s[right : right+wordLen]
			}

			if _, ok := wordsMap[subStr]; !ok {
				left, right = right+wordLen, right+wordLen
				cnt = 0
				winMap = make(map[string]int)
				continue
			}

			if _, ok := winMap[subStr]; !ok {
				winMap[subStr] = 1
			} else {
				winMap[subStr]++
			}
			right += wordLen
			cnt++

			for winMap[subStr] > wordsMap[subStr] {
				winMap[s[left:left+wordLen]]--
				cnt--
				left += wordLen
			}

			if cnt == wordsLen {
				ret = append(ret, left)
			}
		}
	}

	return ret
}

// func findSubstring(s string, words []string) []int {
// 	if len(s) == 0 || len(words) == 0 {
// 		return []int{}
// 	}
// 	var ret = []int{}
// 	for i := 0; i < len(s); i++ {
// 		if checkSubstring(s[i:], words) {
// 			ret = append(ret, i)
// 		}
// 	}
// 	return ret
// }

// func checkSubstring(s string, words []string) bool {
// 	var wordCopy = make([]string, len(words))
// 	copy(wordCopy, words)
// 	fmt.Printf("1wordCopy = %v, s = %s\n", wordCopy, s)
// 	for k, word := range wordCopy {
// 		if len(s) < len(word) {
// 			continue
// 		}
// 		if s[:len(word)] != word {
// 			continue
// 		}
// 		if len(wordCopy) == 1 {
// 			return true
// 		}
// 		if k+1 == len(wordCopy) {
// 			wordCopy = wordCopy[:k]
// 		} else {
// 			wordCopy = append(wordCopy[:k], wordCopy[k+1:]...)
// 		}
// 		s = s[len(word):]
// 		return checkSubstring(s, wordCopy)
// 	}
// 	return false
// }
