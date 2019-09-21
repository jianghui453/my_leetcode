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

//func ladderLength(beginWord string, endWord string, wordList []string) int {
//	wordSet := map[string]bool{}
//	for _, word := range wordList {
//		wordSet[word] = true
//	}
//	if !wordSet[endWord] {
//		return 0
//	}
//	var wl int = len(beginWord)
//	var step int = 0
//	q1 := map[string]bool{beginWord: true}
//	q2 := map[string]bool{endWord: true}
//	for len(q1) > 0 && len(q2) > 0 {
//		step++
//		if len(q1) > len(q2) {
//			q1, q2 = q2, q1
//		}
//		q := map[string]bool{}
//		for w, _ := range q1 {
//			for i := 0; i < wl; i++ {
//				chars := []rune(w)
//				for j := 'a'; j <= 'z'; j++ {
//					chars[i] = j
//					newWord := string(chars)
//					if q2[newWord] {
//						return step + 1
//					}
//					if !wordSet[newWord] {
//						continue
//					}
//					delete(wordSet, newWord)
//					q[newWord] = true
//				}
//			}
//		}
//		q, q1 = q1, q
//	}
//	return 0
//}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if len(wordList) == 0 {
		return 0
	}
	nextWords := make([][]int, len(wordList))
	haveEnd := false
	his := make([]bool, len(wordList))
	words := make([]int, 0)
	for i := range wordList {
		if wordList[i] == endWord {
			haveEnd = true
		}
		for j := i + 1; j < len(wordList); j++ {
			if _isNext(wordList[i], wordList[j]) {
				nextWords[i] = append(nextWords[i], j)
				nextWords[j] = append(nextWords[j], i)
			}
		}
		if _isNext(beginWord, wordList[i]) {
			if wordList[i] == endWord {
				return 2
			}
			words = append(words, i)
			his[i] = true
		}
	}
	if len(words) == 0 || !haveEnd {
		return 0
	}

	min := 2
	for {
		newWords := make([]int, 0)
		for wordsIdx := range words {
            if len(nextWords[words[wordsIdx]]) == 0 {
                continue
            }
            for i := range nextWords[words[wordsIdx]] {
                if his[nextWords[words[wordsIdx]][i]] {
                    continue
                }
                if wordList[nextWords[words[wordsIdx]][i]] == endWord {
                    return min+1
                }
                newWords = append(newWords, nextWords[words[wordsIdx]][i])
                his[nextWords[words[wordsIdx]][i]] = true
            }
		}

		if len(newWords) == 0 {
			return 0
		}
		words = newWords
		min++
	}
}

func _isNext(s1, s2 string) bool {
	cnt := 0
	for i := range s1 {
		if s1[i] != s2[i] {
			cnt++
			if i < len(s1)-1 {
				if s1[:i] == s2[:i] && s1[i+1:] == s2[i+1:] {
					return true
				}
			} else if s1[:i] == s2[:i] {
                return true
			}
			if cnt > 1 {
				return false
			}
		}
	}
	return false
}
