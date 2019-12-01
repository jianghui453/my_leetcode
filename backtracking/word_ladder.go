//给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：
//
//每次转换只能改变一个字母。
//转换过程中的中间单词必须是字典中的单词。
//说明:
//
//如果不存在这样的转换序列，返回 0。
//所有单词具有相同的长度。
//所有单词只由小写字母组成。
//字典中不存在重复的单词。
//你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
//示例 1:
//
//输入:
//beginWord = "hit",
//endWord = "cog",
//wordList = ["hot","dot","dog","lot","log","cog"]
//
//输出: 5
//
//解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
//     返回它的长度 5。
//示例 2:
//
//输入:
//beginWord = "hit"
//endWord = "cog"
//wordList = ["hot","dot","dog","lot","log"]
//
//输出: 0
//
//解释: endWord "cog" 不在字典中，所以无法进行转换。

package backtracking

import (
	// "math"
	// "fmt"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	for i := range wordList {
		if wordList[i] == beginWord {
			if i == len(wordList)-1 {
				wordList = wordList[: i]
			} else {
				wordList = append(wordList[: i], wordList[i+1: ]...)
			}
			break
		}
	}

	r := dfs([]string{beginWord}, wordList, endWord, 1)
	return r
}

func dfs(words, wordList []string, endWord string, cnt int) int {
	l := len(wordList)
	if l == 0 {
		return 0
	}

	wordListHash := make(map[string]bool)
	for i := range words {
		for j := range wordList {
			if isNext(words[i], wordList[j]) {
				if wordList[j] == endWord {
					return cnt+1
				}
				wordListHash[wordList[j]] = true	
			} else if _, ok := wordListHash[wordList[j]]; !ok {
				wordListHash[wordList[j]] = false
			}
		}
	}

	newWords := make([]string, 0)
	newWordList := make([]string, 0)
	for k, v := range wordListHash {
		if v {
			newWords = append(newWords, k)
		} else {
			newWordList = append(newWordList, k)
		}
	}

	r := dfs(newWords, newWordList, endWord, cnt+1)
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
