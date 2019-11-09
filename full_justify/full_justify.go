//给定一个单词数组和一个长度 maxWidth，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。
//
//你应该使用“贪心算法”来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth 个字符。
//
//要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。
//
//文本的最后一行应为左对齐，且单词之间不插入额外的空格。
//
//说明:
//
//单词是指由非空格字符组成的字符序列。
//每个单词的长度大于 0，小于等于 maxWidth。
//输入单词数组 words 至少包含一个单词。
//示例:
//
//输入:
//words = ["This", "is", "an", "example", "of", "text", "justification."]
//maxWidth = 16
//输出:
//[
//   "This    is    an",
//   "example  of text",
//   "justification.  "
//]
//示例 2:
//
//输入:
//words = ["What","must","be","acknowledgment","shall","be"]
//maxWidth = 16
//输出:
//[
//  "What   must   be",
//  "acknowledgment  ",
//  "shall be        "
//]
//解释: 注意最后一行的格式应为 "shall be    " 而不是 "shall     be",
//     因为最后一行应为左对齐，而不是左右两端对齐。
//     第二行同样为左对齐，这是因为这行只包含一个单词。
//示例 3:
//
//输入:
//words = ["Science","is","what","we","understand","well","enough","to","explain",
//         "to","a","computer.","Art","is","everything","else","we","do"]
//maxWidth = 20
//输出:
//[
//  "Science  is  what we",
//  "understand      well",
//  "enough to explain to",
//  "a  computer.  Art is",
//  "everything  else  we",
//  "do                  "
//]

package full_justify

import (
	// "strings"
	// "fmt"
)

func fullJustify(words []string, maxWidth int) []string {
	ret := make([]string, 0)
	l := len(words)
	if l == 0 {
		return ret
	}

	if l == 1 {
		str := words[0]
		for 
		ret = append(ret, words[0] + string(bs))
		return ret
	}

	wordsLen := 0
	i := 0
	for ; i < l; i++ {
		if wordsLen+i+len(words[i]) > maxWidth {
			str := ""
			if i > 1 {
				blanksCnt := maxWidth-wordsLen
				blankStrCnt := i-1
				for j := 0; j < i-1; j++ {
					str += words[j]
					cnt := blanksCnt/(blankStrCnt)
					if blanksCnt%(blankStrCnt) > 0 {
						cnt++
					}
					blanksCnt -= cnt
					blankStrCnt--
					for ; cnt > 0; cnt-- {
						str += " "
					}
				}
				str += words[i-1]
			} else {
				str = words[0]
				for k := 0; k < maxWidth-len(words[0]); k ++ {
					str += " "
				}
			}

			ret = append(ret, str)
			if i < l {
				ret = append(ret, fullJustify(words[i: ], maxWidth)...)
			}
			return ret
		}

		wordsLen += len(words[i])
	}

	str := ""
	for i := 0; i < l-1; i++ {
		str += words[i]
		str += " "
	}
	str += words[l-1]
	for k := 0; k < maxWidth-len(str); k++ {
		str += " "
	}

	ret = append(ret, str)
	return ret
}

// func fullJustify(words []string, maxWidth int) []string {
// 	lenWords := len(words)
// 	if lenWords < 1 {
// 		return []string{}
// 	}
// 	strs := make([]string, 0)
// 	wordsT := make([]string, 0)
// 	lenWordsTCont := 0
// 	for i := 0; i < lenWords; i++ {
// 		lenWord := len(words[i])
// 		lenWordsT := len(wordsT)
// 		if lenWord > maxWidth {
// 			return strs
// 		}
// 		if lenWord+lenWordsTCont+lenWordsT <= maxWidth {
// 			wordsT = append(wordsT, words[i])
// 			lenWordsTCont += lenWord
// 		} else {
// 			sb := strings.Builder{}
// 			if lenWordsT == 1 {
// 				sb.WriteString(wordsT[0])
// 				for j := 0; j < maxWidth-lenWordsTCont; j++ {
// 					sb.WriteString(" ")
// 				}
// 			} else {
// 				cntBlk := maxWidth - lenWordsTCont
// 				cntBlkStr := lenWordsT - 1
// 				blkStrs := make([]string, cntBlkStr)
// 				for j := 0; j < cntBlk; j++ {
// 					idx := j % cntBlkStr
// 					blkStrs[idx] += " "
// 				}
// 				for j := 0; j < len(wordsT); j++ {
// 					sb.WriteString(wordsT[j])
// 					if j < len(wordsT)-1 {
// 						sb.WriteString(blkStrs[j])
// 					}
// 				}
// 			}
// 			strs = append(strs, sb.String())
// 			wordsT = []string{words[i]}
// 			lenWordsTCont = len(words[i])
// 		}
// 	}
// 	if lenWordsTCont > 0 {
// 		sb := strings.Builder{}
// 		lenWordsT := len(wordsT)
// 		for i := 0; i < lenWordsT; i++ {
// 			sb.WriteString(wordsT[i])
// 			if i < lenWordsT-1 {
// 				sb.WriteString(" ")
// 			}
// 		}
// 		for i := 0; i < maxWidth-lenWordsTCont-lenWordsT+1; i++ {
// 			sb.WriteString(" ")
// 		}
// 		strs = append(strs, sb.String())
// 	}
// 	return strs
// }
