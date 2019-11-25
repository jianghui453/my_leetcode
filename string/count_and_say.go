// 报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：

// 1.     1
// 2.     11
// 3.     21
// 4.     1211
// 5.     111221
// 1 被读作  "one 1"  ("一个一") , 即 11。
// 11 被读作 "two 1s" ("两个一"）, 即 21。
// 21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。

// 给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。

// 注意：整数顺序将表示为一个字符串。

//

// 示例 1:

// 输入: 1
// 输出: "1"
// 示例 2:

// 输入: 4
// 输出: "1211"

package string

func countAndSay(n int) string {
	ret := ""
	if n == 0 {
		return ret
	}
	ret = "1"

	for i := 1; i < n; i++ {
		cnt := 1
		_ret := ""
		for j := range ret {
			if j < len(ret)-1 && ret[j] == ret[j+1] {
				cnt++
				continue
			}

			_ret += string(cnt+'0') + string(ret[j])
			cnt = 1
		}

		ret = _ret
	}

	return ret
}

// func countAndSay(n int) string {
// 	if n < 1 {
// 		return ""
// 	}
// 	str := "1"
// 	for ; n > 1; n-- {
// 		var i int
// 		var count int
// 		sb := strings.Builder{}
// 		for _, ch := range str {
// 			if i == 0 {
// 				i = int(ch - '0')
// 				count++
// 				continue
// 			}
// 			if i == int(ch-'0') {
// 				count++
// 				continue
// 			}
// 			fmt.Printf("1count = %d, i = %d, str = %s.\n", count, i, str)
// 			sb.Write([]byte{byte(count + '0')})
// 			sb.Write([]byte{byte(i + '0')})
// 			i = int(ch - '0')
// 			count = 1
// 		}
// 		if count > 0 {
// 			fmt.Printf("2count = %d, i = %d, str = %s.\n", count, i, str)
// 			sb.Write([]byte{byte(count + '0')})
// 			sb.Write([]byte{byte(i + '0')})
// 		}
// 		str = sb.String()
// 	}
// 	return str
// }
