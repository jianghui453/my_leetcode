//一条包含字母 A-Z 的消息通过以下方式进行了编码：
//
//'A' -> 1
//'B' -> 2
//...
//'Z' -> 26
//给定一个只包含数字的非空字符串，请计算解码方法的总数。
//
//示例 1:
//
//输入: "12"
//输出: 2
//解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
//示例 2:
//
//输入: "226"
//输出: 3
//解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。

package dynamic_programming

func numDecodings(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}

	dp := make([]int, l+1)
	for i := 0; i < l; i++ {
		if s[i] > '0' && s[i] <= '9' {
			dp[i] = 1
			for j := i+1; j < l; j++ {
				if s[j] > '0' && s[j] <= '9' {
					dp[j] = dp[j-1]
				}
				if s[j-1] > '0' && (int(s[j-1]-'0')*10 + int(s[j]-'0')) <= 26 {
					if j < 2 {
						dp[j]++
					} else {
						dp[j] += dp[j-2]
					}
				}
				if dp[j] == 0 {
					return 0
				}
			}

			break
		} else {
			return 0
		}
	}
	
	return dp[l-1]
}

// func numDecodings(s string) int {
// 	var r int
// 	var f func(string, int)
// 	f = func(str string, sL int) {
// 		if sL == 0 {
// 			return
// 		}
// 		if sL == 1 {
// 			if str[0]-'0' > 0 {
// 				r++
// 			}
// 			return
// 		}
// 		if str[0]-'0' > 0 {
// 			f(str[1:], sL-1)
// 			n := 10*(str[0]-'0') + (str[1] - '0')
// 			if n > 0 && n <= 26 {
// 				if sL == 2 {
// 					r++
// 				} else {
// 					f(str[2:], sL-2)
// 				}
// 			}
// 		}
// 	}
// 	f(s, len(s))
// 	return r
// }

//func numDecodings(s string) int {
//	l := len(s)
//	if l == 0 {
//	    return 0
//    }
//	if l == 1 {
//	    if s[0]-'0' > 0 {
//	        return 1
//        }
//	    return 0
//    }
//	dp := make([]int, l)
//    if s[0]-'0' > 0 {
//        dp[0] = 1
//    } else {
//        return 0
//    }
//	fmt.Printf("dp=%v\n", dp)
//	if s[1]-'0' > 0 {
//	    dp[1] = dp[0]
//    }
//    fmt.Printf("dp=%v\n", dp)
//    n := (s[1]-'0') + (s[0] - '0')*10
//    fmt.Printf("n=%d\n", n)
//    if n > 0 && n <= 26 {
//        dp[1] += 1
//    }
//    fmt.Printf("dp=%v\n", dp)
//	for i := 2; i < l; i++ {
//	    if s[i]-'0'>0 {
//	        dp[i] = dp[i-1]
//        }
//	    if s[i-1]-'0' > 0 {
//            n := (s[i-1]-'0')*10 + (s[i] - '0')
//            if n > 0 && n <= 26 {
//                dp[i] += dp[i-2]
//            }
//        }
//	}
//	return dp[l-1]
//}
