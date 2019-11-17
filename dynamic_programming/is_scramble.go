//给定一个字符串 s1，我们可以把它递归地分割成两个非空子字符串，从而将其表示为二叉树。
//
//下图是字符串 s1 = "great" 的一种可能的表示形式。
//
//    great
//   /    \
//  gr    eat
// / \    /  \
//g   r  e   at
//           / \
//          a   t
//在扰乱这个字符串的过程中，我们可以挑选任何一个非叶节点，然后交换它的两个子节点。
//
//例如，如果我们挑选非叶节点 "gr" ，交换它的两个子节点，将会产生扰乱字符串 "rgeat" 。
//
//    rgeat
//   /    \
//  rg    eat
// / \    /  \
//r   g  e   at
//           / \
//          a   t
//我们将 "rgeat” 称作 "great" 的一个扰乱字符串。
//
//同样地，如果我们继续交换节点 "eat" 和 "at" 的子节点，将会产生另一个新的扰乱字符串 "rgtae" 。
//
//    rgtae
//   /    \
//  rg    tae
// / \    /  \
//r   g  ta  e
//       / \
//      t   a
//我们将 "rgtae” 称作 "great" 的一个扰乱字符串。
//
//给出两个长度相等的字符串 s1 和 s2，判断 s2 是否是 s1 的扰乱字符串。
//
//示例 1:
//
//输入: s1 = "great", s2 = "rgeat"
//输出: true
//示例 2:
//
//输入: s1 = "abcde", s2 = "caebd"
//输出: false

package dynamic_programming

import (
	// "fmt"
)

func isScramble(s1 string, s2 string) bool {
	l := len(s1)
	if len(s2) != l {
		return false
	}
	if l == 0 {
		return true
	}

	dp := make([][][]bool, l)
	for i := 0; i < l; i++ {
		dp[i] = make([][]bool, l)
		
		for j := 0; j < l; j++ {
			dp[i][j] = make([]bool, l+1)
			dp[i][j][0] = true
			if s1[i] == s2[j] {
				dp[i][j][1] = true
			}
		}
		if l == 1 {
			return s1 == s2
		}

	for i := l-1; i >= 0; i-- {
		for j := l-1; j >= 0; j-- {
			_l := l-i
			if l-j < l-i {
				_l = l-j
			}

			for k1 := 1; k1 < _l; k1++ {
				for k2 := 1; k2 <= _l-k1; k2++ {
					if (j+k1 < l && i+k2 < l && dp[i][j+k1][k2] && dp[i+k2][j][k1]) ||
							(i+k1 < l && j+k1 < l && dp[i][j][k1] && dp[i+k1][j+k1][k2]) {
						dp[i][j][k1+k2] = true
					}
				}
			}
		}

	return dp[0][0][l]
}

/** [
		[
			[true true false false false] 
			[true false false false false] 
			[true false false false false] 
			[true false false false false]
		] 
		[
			[true false false false false] 
			[true false false true false] 
			[true false false false false] 
			[true true false false false]
		] 
		[
			[true false false false false] 
			[true true true false false] 
			[true false false false false] 
			[true false false false false]
		] 
		[
			[true false false false false] 
			[true false false false false] 
			[true true false false false] 
			[true false false false false]
		]
	]
*/
//func isScramble(s1 string, s2 string) bool {
//	len1 := len(s1)
//	len2 := len(s2)
//	if len1 != len2 {
//		return false
//	}
//	if len1 < 2 {
//		return s1 == s2
//	}
//	memo := make(map[string]bool)
//	return recursion(s1, s2, len1, &memo)
//}
//
//func recursion(s1, s2 string, lenS int, memo *map[string]bool) bool {
//	if lenS == 1 {
//		return s1 == s2
//	}
//	key := s1+s2
//	if v, ok := (*memo)[key]; ok {
//	    return v
//    }
//	for i := 1; i < lenS; i++ {
//		if (recursion(s1[:i], s2[:i], i, memo) && recursion(s1[i:], s2[i:], lenS-i, memo)) ||
//			(recursion(s1[:i], s2[lenS-i:], i, memo) && recursion(s1[i:], s2[:lenS-i], lenS-i, memo)) {
//            (*memo)[key] = true
//			return true
//		}
//	}
//    (*memo)[key] = false
//	return false
//}
