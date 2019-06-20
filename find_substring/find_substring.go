package find_sub_string

import (
	"fmt"
)

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}
	wordsMap := make(map[string]int)
	for _, word := range words {
		if _, ok := wordsMap[word]; !ok {
			wordsMap[word] = 1
		} else {
			wordsMap[word]++
		}
	}
	wordLen := len(words[0])
	var ret = []int{}
	Loop:
	for i := 0; i <= len(s) - wordLen * len(words); i ++ {
		sMap := make(map[string]int)
		for j := i; j < i + wordLen * len(words); j += wordLen {
			var str string
			if j + wordLen == len(s) {
				str = s[j:]
			} else {
				str = s[j: j + wordLen]
			}
			if _, ok := wordsMap[str]; !ok {
				continue Loop
			} else if _, ok := sMap[str]; !ok {
				sMap[str] = 1
			} else {
				sMap[str]++
			}
			if sMap[str] > wordsMap[str] {
				continue Loop
			}
		}
		ret = append(ret, i)
	}
	return ret
}

func findSubstringV1(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}
	var ret = []int{}
	for i := 0; i < len(s); i ++ {
		if checkSubstring(s[i:], words) {
			ret = append(ret, i)
		}
	}
	return ret
}

func checkSubstring(s string, words []string) bool {
	var wordCopy = make([]string, len(words))
	copy(wordCopy, words)
	fmt.Printf("1wordCopy = %v, s = %s\n", wordCopy, s)
	for k, word := range wordCopy {
		if len(s) < len(word) {
			continue
		}
		if s[:len(word)] != word {
			continue
		}
		if len(wordCopy) == 1 {
			return true
		}
		if k + 1 == len(wordCopy) {
			wordCopy = wordCopy[:k]
		} else {
			wordCopy = append(wordCopy[:k], wordCopy[k + 1:]...)
		}
		s = s[len(word):]
		return checkSubstring(s, wordCopy)
	}
	return false;
}