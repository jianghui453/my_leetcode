//给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串。
//
//示例：
//
//输入: S = "ADOBECODEBANC", T = "ABC"
//输出: "BANC"
//说明：
//
//如果 S 中不存这样的子串，则返回空字符串 ""。
//如果 S 中存在这样的子串，我们保证它是唯一的答案。

package sliding_window

import (
	"math"
	// "fmt"
)

func minWindow(s string, t string) string {
	ls := len(s)
	lt := len(t)

	if ls < lt || lt == 0 {
		return ""
	}

	hashS := make([]int, 256)
	hashT := make([]int, 256)

	cntT := 0
	for i := 0; i < lt; i++ {
		idx := int(t[i])
		hashT[idx]++
		if hashT[idx] == 1 {
			cntT++
		}
	}

	strRet := ""
	strLen := math.MaxInt64
	cnt := 0
	left, right := 0, 0

	for ; right < ls; right++ {
		// fmt.Printf("left=%d right=%d cnt=%d cntT=%d\n", left, right, cnt, cntT)
		idx := s[right]

		if hashT[idx] == 0 {
			continue
		}

		hashS[idx]++
		if hashS[idx] == hashT[idx] {
			cnt++
		}

		if cnt < cntT {
			continue
		}
		
		for ; left <= right; left++ {
			// fmt.Printf("left=%d\n", left)
			if right-left+1 < strLen {
				if right < ls-1 {
					strRet = s[left: right+1]
				} else {
					strRet = s[left: ]
				}
				strLen = right-left+1
			}

			idx := int(s[left])
			if hashT[idx] == 0 {
				continue
			}

			hashS[idx]--
			if hashS[idx] == hashT[idx]-1 {
				cnt--
				left++
				break
			}
		}
	}

	return strRet
}

// func minWindow(s string, t string) string {
// 	countList := make([]int32, 256)
// 	for _, c := range t {
// 		countList[c]++
// 	}

// 	cnt := 0
// 	minLen := len(s)
// 	var res string

// 	for i, j := 0, 0; j < len(s); j++ {
// 		if countList[s[j]] >= 1 {
// 			cnt++
// 		}
// 		countList[s[j]]--
// 		for cnt == len(t) {
// 			if j-i+1 <= minLen {
// 				minLen = j - i + 1
// 				res = s[i : j+1]
// 			}
// 			countList[s[i]]++
// 			if countList[s[i]] > 0 {
// 				cnt--
// 			}
// 			i++
// 		}
// 	}
// 	return res
// }

//func minWindow(s string, t string) string {
//    lenS := len(s)
//    lenT := len(t)
//    if lenS*lenT == 0 || lenS < lenT {
//        return ""
//    }
//    rcdT := make(map[uint8]int)
//    for i := 0; i < lenT; i++ {
//        k := t[i]
//        if _, ok := rcdT[k]; ok {
//            rcdT[k]++
//        } else {
//            rcdT[k] = 1
//        }
//    }
//    //fmt.Printf("rcdT=%v\n", rcdT)
//    left := 0
//    right := 0
//    minStr := ""
//    cntRcdS := 0
//    rcdS := make(map[uint8]int)
//    if _, ok := rcdT[s[0]]; ok {
//        rcdS[s[0]] = 1
//        cntRcdS ++
//    }
//loop:
//    for {
//        //fmt.Printf("rcdS=%v \n left=%d \nright=%d\n", rcdS, left, right)
//        if cntRcdS < lenT {
//            if right < lenS - 1 {
//                goto mvRight
//            }
//            break loop
//        }
//        for k, v := range rcdT {
//            if _, ok := rcdS[k]; !ok || rcdS[k] < v {
//                if right < lenS - 1 {
//                    goto mvRight
//                }
//                break loop
//            }
//        }
//        if minStr == "" || right-left+1 < len(minStr) {
//            if right < lenS-1 {
//                minStr = s[left: right+1]
//            } else {
//                minStr = s[left: ]
//            }
//        }
//        if cntRcdS >= lenT {
//            goto mvLeft
//        } else {
//            if right < lenS - 1 {
//                goto mvRight
//            }
//            break loop
//        }
//    mvRight:
//        right ++
//        if _, ok := rcdT[s[right]]; !ok {
//            continue
//        }
//        if _, ok := rcdS[s[right]]; ok {
//            rcdS[s[right]] ++
//        } else {
//            rcdS[s[right]] = 1
//        }
//        cntRcdS ++
//        //fmt.Printf("right\n")
//        continue
//    mvLeft:
//        if _, ok := rcdS[s[left]]; ok {
//            rcdS[s[left]] --
//            cntRcdS --
//        }
//        left ++
//        //fmt.Printf("left\n")
//    }
//    return minStr
//}

//func minWindow(s string, t string) string {
//    lenS := len(s)
//    lenT := len(t)
//    if lenS*lenT == 0 || lenS < lenT {
//        return ""
//    }
//    rcdT := make(map[uint8]int)
//    for i := 0; i < lenT; i++ {
//        k := t[i]
//        if _, ok := rcdT[k]; ok {
//            rcdT[k]++
//        } else {
//            rcdT[k] = 1
//        }
//    }
//    for winWidth := lenT; winWidth <= lenS; winWidth ++ {
//        rcdS := make(map[uint8]int)
//        for i := 0; i < winWidth; i ++ {
//            k := s[i]
//            if _, ok := rcdS[k]; ok {
//                rcdS[k]++
//            } else {
//                rcdS[k] = 1
//            }
//        }
//        loop:
//        for winStart := 0; winStart+winWidth-1 < lenS; winStart ++ {
//            if winStart > 0 {
//                rcdS[s[winStart-1]]--
//                if _, ok := rcdS[s[winStart+winWidth-1]]; ok {
//                    rcdS[s[winStart+winWidth-1]]++
//                } else {
//                    rcdS[s[winStart+winWidth-1]] = 1
//                }
//                if _, ok := rcdT[s[winStart+winWidth-1]]; !ok {
//                    continue loop
//                }
//            }
//            for k, v := range rcdT {
//                if _, ok := rcdS[k]; !ok || rcdS[k] < v {
//                    continue loop
//                }
//            }
//            if winStart+winWidth < lenS {
//                return s[winStart: winStart+winWidth]
//            } else {
//                return s[winStart: ]
//            }
//        }
//    }
//    return ""
//}
