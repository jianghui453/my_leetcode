//给定两个单词（beginWord 和 endWord）和一个字典 wordList，找出所有从 beginWord 到 endWord 的最短转换序列。转换需遵循如下规则：
//
//每次转换只能改变一个字母。
//转换过程中的中间单词必须是字典中的单词。
//说明:
//
//如果不存在这样的转换序列，返回一个空列表。
//所有单词具有相同的长度。
//所有单词只由小写字母组成。
//字典中不存在重复的单词。
//你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
//示例 1:
//
//输入:
//beginWord = "hit",
//endWord = "cog",
//wordList = ["hot","dot","dog","lot","log","cog"]
//
//输出:
//[
//  ["hit","hot","dot","dog","cog"],
//  ["hit","hot","lot","log","cog"]
//]
//示例 2:
//
//输入:
//beginWord = "hit"
//endWord = "cog"
//wordList = ["hot","dot","dog","lot","log"]
//
//输出: []
//
//解释: endWord "cog" 不在字典中，所以不存在符合要求的转换序列。

package backtracking

import (
)

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	words := [][]string{{beginWord}}
	for i := range wordList {
		if wordList[i] == beginWord {
			var r [][]string
			if i == len(wordList)-1 {
				r = bfs(words, wordList[: len(wordList)-1], endWord)
			} else {
				r = bfs(words, append(wordList[: i], wordList[i+1: ]...), endWord)
			}
			return r
		}
	}

	r := bfs(words, wordList, endWord)
	return r
}

func bfs(words [][]string, wordList []string, endWord string) [][]string {
	if len(wordList) == 0 {
		return [][]string{}
	}
	if len(words) == 0 {
		return [][]string{}
	}

	newWords := make([][]string, 0)
	ret := make([][]string, 0)
	memo := make(map[string]bool)

	l := len(words[0])
	for i := range words {
		word := words[i][l-1]

		for j := range wordList {
			if isNext(word, wordList[j]) {
				wordsI := make([]string, len(words[i]))
				copy(wordsI, words[i])
				if wordList[j] == endWord {
					ret = append(ret, append(wordsI, wordList[j]))
				} else {
					newWords = append(newWords, append(wordsI, wordList[j]))
				}
				memo[wordList[j]] = true
			} else if _, ok := memo[wordList[j]]; !ok {
				memo[wordList[j]] = false
			}
		}
	}
	if len(ret) > 0 {
		return ret
	}

	newWordList := make([]string, 0)
	for k, v := range memo {
		if !v {
			newWordList = append(newWordList, k)
		}
	}

	r := bfs(newWords, newWordList, endWord)
	return r
}

func isNext(s1, s2 string) bool {
	l := len(s1)
	if l == 1 {
		return true
	}

	left, right := 0, l-1
	for left < right {
		if s1[left] != s2[left] && s1[right] != s2[right] {
			return false
		}
		if s1[left] == s2[left] {
			left++
		}
		if s1[right] == s2[right] {
			right--
		}
	}

	return true
}
