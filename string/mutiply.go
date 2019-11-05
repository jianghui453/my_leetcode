// Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.

// Example 1:

// Input: num1 = "2", num2 = "3"
// Output: "6"
// Example 2:

// Input: num1 = "123", num2 = "456"
// Output: "56088"
// Note:

// The length of both num1 and num2 is < 110.
// Both num1 and num2 contain only digits 0-9.
// Both num1 and num2 do not contain any leading zero, except the number 0 itself.
// You must not use any built-in BigInteger library or convert the inputs to integer directly.

package string

import (
	// "fmt"
)

func multiply(num1 string, num2 string) string {
	num1Len := len(num1)
	num2Len := len(num2)

	if num1Len * num2Len == 0 {
		return ""
	}
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	ret := make([]int, 0)
	for i := num1Len-1; i >= 0; i-- {
		for j := num2Len-1; j >= 0; j-- {
			product := int(num1[i]-'0') * int(num2[j]-'0')
			
			k := num1Len+num2Len-2-i-j
			for product > 0 {
				if len(ret) <= k {
					ret = append(ret, make([]int, k-len(ret)+1)...)
				}
				
				product += ret[k]
				ret[k] = product%10
				product = product/10
				k++
			}
		}
	}

	retLen := len(ret)
	str := make([]byte, retLen)
	for i := 0; i < retLen; i++ {
		str[retLen-1-i] = byte(ret[i] + '0')
	}

	return string(str)
}

// func multiply(num1 string, num2 string) string {
// 	if num1 == "0" || num2 == "0" {
// 		return "0"
// 	}

// 	// temp 选用为[]int，而不是[]byte，是因为
// 	// Go中，byte的基础结构是uint8，最大值为255。
// 	// 不考虑进位的话，temp会溢出
// 	temp := make([]int, len(num1)+len(num2))
// 	for i := 0; i < len(num1); i++ {
// 		for j := 0; j < len(num2); j++ {
// 			// 每个结果位上的数字，等于两个字符不同索引和等于该索引的位置上数字乘积的和
// 			// 如结果中的第3位，由第（1、2）、（2、1）位上乘积的和组成
// 			temp[i+j+1] += int(num1[i]-'0') * int(num2[j]-'0')
// 		}
// 	}

// 	// 进位：每个位上转为个数
// 	for i := len(temp) - 1; i > 0; i-- {
// 		temp[i-1] += temp[i] / 10
// 		temp[i] = temp[i] % 10
// 	}

// 	// 去掉首位0
// 	if temp[0] == 0 {
// 		temp = temp[1:]
// 	}

// 	// 转换结果，int转byte转string
// 	res := make([]byte, len(temp))
// 	for i := 0; i < len(temp); i++ {
// 		res[i] = '0' + byte(temp[i])
// 	}

// 	return string(res)
// }

// func multiply(num1, num2 string) string {
// 	if len(num1) == 0 || len(num2) == 0 {
// 		return "0"
// 	}
// 	if num1 == "0" || num2 == "0" {
// 		return "0"
// 	}
// 	ret := []int{}
// 	for i1 := len(num1) - 1; i1 >= 0; i1-- {
// 		if num1[i1] == '0' {
// 			continue
// 		}
// 		for i2 := len(num2) - 1; i2 >= 0; i2-- {
// 			if num1[i2] == '0' {
// 				continue
// 			}
// 			val := int(num1[i1]-'0') * int(num2[i2]-'0')
// 			idx := len(num1) + len(num2) - (i1 + i2) - 2
// 			for {
// 				if len(ret) <= idx {
// 					ret = append(ret, make([]int, idx-len(ret)+1)...)
// 				}
// 				if val+ret[idx] < 10 {
// 					ret[idx] += val
// 					break
// 				}
// 				val += ret[idx]
// 				ret[idx] = val % 10
// 				if idx == len(ret)-1 {
// 					ret = append(ret, val/10)
// 					break
// 				}
// 				val = val / 10
// 				idx++
// 			}
// 		}
// 	}
// 	sb := strings.Builder{}
// 	for i := len(ret) - 1; i >= 0; i-- {
// 		sb.WriteByte(byte(ret[i]) + '0')
// 	}
// 	return sb.String()
// }
