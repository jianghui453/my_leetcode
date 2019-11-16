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

func isScramble(s1 string, s2 string) bool {
	memo := make(map[string]map[string]bool)
	var f func(str1, str2 string) bool

	f = func(str1, str2 string) bool {
		l := len(s1)
		if len(s2) != l {
			return false
		}
		if l == 1 {
			return s1 == s2
		}

		if _, ok := memo[str1]; ok {
			if _, _ok := memo[str1][str2]; ok {
				return memo[str1][str2]
			}
		}

		if l == 2 {
			s3 := s2[1] + s2[0]
			ret := s1 == s2 || s1 == s3
			memo[str1][str2] = ret
			return ret
		}

		memo[str1][str2] = false
		for i := 1; i <= l/2; i++ {
			if isScramble(s1[: i], s2[: i]) || isScramble(s[l-i: ], s2[: i]) {
				memo[str1][str2] = true
				break
			}
		}

		return memo[str1][str2]
	}
	
	r := f(s1, s2)
	return r
}

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
