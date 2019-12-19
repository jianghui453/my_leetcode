// 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

// 注意：

// num1 和num2 的长度都小于 5100.
// num1 和num2 都只包含数字 0-9.
// num1 和num2 都不包含任何前导零。
// 你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。

package string

func addStrings(num1 string, num2 string) string {
	l1, l2 := len(num1), len(num2)
	if l1 == 0 {
		return num2
	}
	if l2 == 0 {
		return num1
	}

	if l1 < l2 {
		l1, l2, num1, num2 = l2, l1, num2, num1
	}

	ret := make([]byte, l1 + 1)
	i := 1
	carry := 0
	
	for i <= l1 && i <= l2 {
		sum := int(num1[l1-i]-'0' + num2[l2-i]-'0') + carry
		ret[l1+1-i] = byte(sum%10+'0')
		carry, i = sum/10, i+1
	}

	for i <= l1 {
		sum := int(num1[l1-i]-'0') + carry
		ret[l1+1-i] = byte(sum%10+'0')
		carry, i = sum/10, i+1
	}

	if carry > 0 {
		ret[0] = byte(carry+'0')
	}

	if ret[0] == 0 {
		ret = ret[1: ]
	}

	return string(ret)
}