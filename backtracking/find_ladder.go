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
	"strconv"
	"strings"
)

//func prevHandle(wordList []string) map[string][]string {
//    regex2Word := make(map[string][]string, 0)
//    for index := range wordList {
//        word := wordList[index]
//        for i := 0; i < len(word); i++ {
//            tmp := word[:i] + "*" + word[i+1:]
//            _, ok := regex2Word[tmp]
//            if !ok {
//                regex2Word[tmp] = make([]string, 0)
//            }
//            regex2Word[tmp] = append(regex2Word[tmp], word)
//        }
//    }
//    return regex2Word
//}
//
//func dfs(path map[string][]string, beginWord string, endWord string, tmpPath []string, result *[][]string) {
//    if beginWord == endWord {
//        newPath := make([]string, len(tmpPath))
//        copy(newPath, tmpPath)
//        *result = append(*result, newPath)
//        return
//    }
//
//    prevWords := path[endWord]
//    for i := range prevWords {
//        prevWord := prevWords[i]
//        dfs(path, beginWord, prevWord, append(append([]string{}, prevWord), tmpPath...), result)
//    }
//}
//
//func findLadders(beginWord string, endWord string, wordList []string) [][]string {
//    regex2Word := prevHandle(wordList)
//    queue := make([]string, 0)
//    queue = append(queue, beginWord)
//
//    results := make([][]string, 0)
//    path := make(map[string][]string, 0)
//    flag := make(map[string]bool, 0)
//
//    addPath := func(prevWord string, word string) {
//        if _, ok := path[word]; !ok {
//            path[word] = make([]string, 0)
//        }
//        needAdd := true
//        for _, old := range path[word] {
//            if old == prevWord {
//                needAdd = false
//            }
//        }
//        if needAdd {
//            path[word] = append(path[word], prevWord)
//        }
//    }
//
//    isHaveResult := false
//
//    for {
//        nextQueue := make([]string, 0)
//        for len(queue) != 0 {
//            word := queue[0]
//            queue = queue[1:]
//            for i := 0; i < len(word); i++ {
//                reg := word[:i] + "*" + word[i+1:]
//                list, ok := regex2Word[reg]
//                if !ok {
//                    continue
//                }
//                for _, tmpNextWord := range list {
//                    if tmpNextWord == word {
//                        continue
//                    }
//                    if fl, ok := flag[tmpNextWord]; ok && fl {
//                        continue
//                    }
//                    addPath(word, tmpNextWord)
//                    if tmpNextWord == endWord {
//                        isHaveResult = true
//                        break
//                    } else {
//                        nextQueue = append(nextQueue, tmpNextWord)
//                    }
//                }
//            }
//        }
//
//        queue = nextQueue
//        if len(queue) == 0 {
//            break
//        }
//        for _, word := range queue {
//            flag[word] = true
//        }
//    }
//
//    newResult := make([][]string, 0)
//    if isHaveResult {
//        tmp := make([]string, 1)
//        tmp[0] = endWord
//        dfs(path, beginWord, endWord, tmp, &results)
//        min := math.MaxInt32
//        for _, result := range results {
//            if len(result) < min {
//                min = len(result)
//            }
//        }
//        for _, result := range results {
//            if len(result) == min {
//                newResult = append(newResult, result)
//            }
//        }
//
//    }
//    return newResult
//}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wl := len(wordList)
	ret := make([][]string, 0)
	if wl == 0 {
		return ret
	}
	matrix := make([][]bool, wl)
	for i := 0; i < wl; i++ {
		matrix[i] = make([]bool, wl)
	}
	beginWords := make([][]int, 0)
	beginWordsLen := 0
	beginWordsEnd := make([]bool, wl)
	beginWordsUsage := make([]bool, wl)

	endWordIdx := -1
	for i := 0; i < wl; i++ {
		// 记录起始单词可以转换的单词
		if isNext(beginWord, wordList[i]) {
			beginWords = append(beginWords, []int{i})
			// 可以直接转换到结束单词
			if wordList[i] == endWord {
				return [][]string{{beginWord, endWord}}
			}
			beginWordsLen++
			beginWordsEnd[i] = true
			beginWordsUsage[i] = true
		}
		for j := i + 1; j < wl; j++ {
			if isNext(wordList[i], wordList[j]) {
				matrix[i][j] = true
				matrix[j][i] = true
			}
		}
		if wordList[i] == endWord {
			endWordIdx = i
		}
	}
	// wordList不包含结束单词
	if endWordIdx == -1 {
		return ret
	}
	endWords := make([][]int, 0)
	endWordsLen := 0
	endWordsUsage := make([]bool, wl)
	endWordsEnd := make([]bool, wl)
	endWordsUsage[endWordIdx] = true
	// 从结束单词开始查找
	for i := 0; i < wl; i++ {
		if matrix[endWordIdx][i] {
			endWords = append(endWords, []int{i})
			endWordsLen++
			endWordsUsage[i] = true
			endWordsEnd[i] = true
		}
	}

	level := 1
	for beginWordsLen > 0 && endWordsLen > 0 {
		//fmt.Printf("beginWords=%v \nendWords=%v\n", beginWords, endWords)
		// 检查查找是否有连接了
		for j := 0; j < wl; j++ {
			if beginWordsUsage[j] && endWordsUsage[j] {
				ret1 := make([][]string, 0)
				ret1Len := 0
				for k := 0; k < beginWordsLen; k++ {
					rpt := make(map[string]bool)
					for x := 0; x < endWordsLen; x++ {
						if level > 1 && beginWords[k][level-1] == endWords[x][level-2] {
							retItem := make([]string, level*2)
							retItem[0] = beginWord
							for y := 1; y <= level; y++ {
								retItem[y] = wordList[beginWords[k][y-1]]
							}
							rptKeySb := strings.Builder{}
							for y := 0; y < level-2; y++ {
								rptKeySb.WriteString(strconv.Itoa(endWords[x][level-3-y]))
								retItem[level+y+1] = wordList[endWords[x][level-3-y]]
							}
							rptKey := rptKeySb.String()
							if _, ok := rpt[rptKey]; !ok {
								retItem[2*level-1] = endWord
								ret1 = append(ret1, retItem)
								ret1Len++
								rpt[rptKey] = true
							}
						} else if ret1Len == 0 && beginWords[k][level-1] == endWords[x][level-1] {
							retItem := make([]string, level*2+1)
							retItem[0] = beginWord
							for y := 1; y <= level; y++ {
								retItem[y] = wordList[beginWords[k][y-1]]
							}
							for y := 0; y < level-1; y++ {
								retItem[level+y+1] = wordList[endWords[x][level-2-y]]
							}
							retItem[2*level] = endWord
							ret = append(ret, retItem)
						}
					}
				}
				if ret1Len > 0 {
					return ret1
				}
				return ret
			}
		}
		if 2*level > wl {
			break
		}

		// 继续查找
		newBeginWords := make([][]int, 0)
		newBeginWordsLen := 0
		newBeginWordsEnd := make([]bool, wl)
		// 查找从起始单词开始的单词
		for i := 0; i < beginWordsLen; i++ {
			for j := 0; j < wl; j++ {
				if (!beginWordsUsage[j] || newBeginWordsEnd[j]) && matrix[beginWords[i][level-1]][j] {
					t := make([]int, level)
					copy(t, beginWords[i])
					newBeginWords = append(newBeginWords, append(t, j))
					newBeginWordsLen++
					newBeginWordsEnd[j] = true
					beginWordsUsage[j] = true
				}
			}
		}
		beginWords = newBeginWords
		beginWordsLen = newBeginWordsLen
		beginWordsEnd = newBeginWordsEnd

		newEndWords := make([][]int, 0)
		newEndWordsLen := 0
		newEndWordsEnd := make([]bool, wl)
		for i := 0; i < endWordsLen; i++ {
			for j := 0; j < wl; j++ {
				if (!endWordsUsage[j] || newEndWordsEnd[j]) && matrix[endWords[i][level-1]][j] {
					t := make([]int, level)
					copy(t, endWords[i])
					newEndWords = append(newEndWords, append(t, j))
					newEndWordsLen++
					newEndWordsEnd[j] = true
					endWordsUsage[j] = true
				}
			}
		}
		endWords = newEndWords
		endWordsLen = newEndWordsLen
		endWordsEnd = newEndWordsEnd
		level++
	}
	//fmt.Printf("beginWords=%v \nendWords=%v\n", beginWords, endWords)
	return ret
}

func isNext(s1, s2 string) bool {
	for i := range s1 {
		if i < len(s1)-1 {
			_s1 := s1[:i] + "*" + s1[i+1:]
			_s2 := s2[:i] + "*" + s2[i+1:]
			if _s1 == _s2 {
				return true
			}
		} else {
			return s1[:i] == s2[:i]
		}
	}
	return false
}
