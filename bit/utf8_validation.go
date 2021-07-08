// 这是 UTF-8 编码的工作方式：

//    Char. number range  |        UTF-8 octet sequence
//       (hexadecimal)    |              (binary)
//    --------------------+---------------------------------------------
//    0000 0000-0000 007F | 0xxxxxxx
//    0000 0080-0000 07FF | 110xxxxx 10xxxxxx
//    0000 0800-0000 FFFF | 1110xxxx 10xxxxxx 10xxxxxx
//    0001 0000-0010 FFFF | 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
// 给定一个表示数据的整数数组，返回它是否为有效的 utf-8 编码。

// 注意:
// 输入是整数数组。只有每个整数的最低 8 个有效位用来存储数据。这意味着每个整数只表示 1 字节的数据。

// 示例 1:

// data = [197, 130, 1], 表示 8 位的序列: 11000101 10000010 00000001.

// 返回 true 。
// 这是有效的 utf-8 编码，为一个2字节字符，跟着一个1字节字符。
// 示例 2:

// data = [235, 140, 4], 表示 8 位的序列: 11101011 10001100 00000100.

// 返回 false 。
// 前 3 位都是 1 ，第 4 位为 0 表示它是一个3字节字符。
// 下一个字节是开头为 10 的延续字节，这是正确的。
// 但第二个延续字节不以 10 开头，所以是不符合规则的。

package bit

import (
// "fmt"
)

func validUtf8(data []int) bool {
	l := len(data)
	for i := 0; i < l; i++ {
		if data[i]&(1<<7) == 0 {
			continue
		}

		cnt := getCnt(data[i])
		if cnt == -1 || i+cnt >= l {
			return false
		}

		if !_validUtf8(i+cnt+1, i+1, data) {
			return false
		}

		i += cnt
	}

	return true
}

func getCnt(num int) int {
	bit := 1<<7 | 1<<6
	if num&bit == bit && num&(1<<5) == 0 {
		return 1
	}

	bit |= (1 << 5)
	if num&bit == bit && num&(1<<4) == 0 {
		return 2
	}

	bit |= (1 << 4)
	if num&bit == bit && num&(1<<3) == 0 {
		return 3
	}

	return -1
}

func _validUtf8(max, i int, data []int) bool {
	for ; i < max; i++ {
		if data[i]&(1<<7) != (1<<7) || data[i]&(1<<6) != 0 {
			return false
		}
	}
	return true
}
