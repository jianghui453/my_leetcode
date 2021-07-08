// 给定一个整数 n ，你需要找到与它最近的回文数（不包括自身）。

// “最近的”定义为两个整数差的绝对值最小。

// 示例 1:

// 输入: "123"
// 输出: "121"
// 注意:

// n 是由字符串表示的正整数，其长度不超过18。
// 如果有多个结果，返回最小的那个。

package string

import (
	"fmt"
)

func nearestPalindromic(n string) string {
	var (
		l = len(n)
		m = []byte(n)
		_m = make([]byte, l)
		vb, vs string
	)

	if l == 0 {
		return ""
	}

	if l == 1 {
		if n[0] == '0' {
			return ""
		}
		return string(byte(n[0] - 1))
	}

	// 将前半段镜像复制到后半段
	for i := 0; i < l/2; i++ {
		m[l-1-i] = m[i]
	}

	copy(_m, m)
	if compare(string(m), n) >= 0 {
		if string(m) != n {
			vb = string(m)
		}
		for i := (l-1) / 2; i >= 0; i-- {
			if (i > 0 && m[i] == '0') || (i == 0 && m[i] == '1') {
				m[i], m[l-1-i] = '9', '9'
				continue
			}
			// fmt.Println("0 i =", i, "m[i] =", string(m[i]), "m[l-1-i] =", string(m[l-1-i]))
			m[i], m[l-1-i] = m[i]-1, m[i]-1

			// fmt.Println("1 i =", i, "m[i] =", string(m[i]), "m[l-1-i] =", string(m[l-1-i]))
			if compare(string(m), n) < 0 {
				vs = string(m)
				break
			}
			m[i], m[l-1-i] = '9', '9'
		}
	}

	fmt.Println("-1 vs =", vs, "vb =", vb)
	copy(m, _m)
	if compare(string(m), n) <= 0 {
		if string(m) != n && (vs == "" || (string(m) != n && compare(string(m), vs) > 0)) {
			vs = string(m)
		}

		for i := (l-1) / 2; i >= 0; i-- {
			if m[i] == '9' {
				m[i], m[l-1-i] = '0', '0'
				continue
			}
			m[i], m[l-1-i] = m[i]+1, m[i]+1

			if string(m) == n {
				m[i], m[l-1-i] = '0', '0'
				continue
			}
			if compare(string(m), n) > 0 && (vb == "" || compare(string(m), vb) > 0) {
				vb = string(m)
			}
			break
		}
	}

	fmt.Println("0 vs =", vs, "vb =", vb)
	mstr := string(_m)
	if vs == "" {
		vs = defVs(mstr)
	}
	if vb == "" {
		vb = defVb(mstr)
	}
	fmt.Println("1 vs =", vs, "vb =", vb)
	if compare(minus(vb, n), minus(n, vs)) < 0 {
		return vb
	}
	return vs
}

func compare(m, n string) int {
	if len(m) > len(n) {
		return  1
	}
	if len(m) < len(n) {
		return -1
	}
	if m == n {
		return 0
	}
	for i := 0; i < len(m); i++ {
		if m[i] > n[i] {
			return 1
		} else if m[i] < n[i] {
			break
		}
	}
	return -1
}

func defVb(m string) string {
	ret := make([]byte, len(m)+1)
	ret[0], ret[len(m)] = '1', '1'
	for i := 1; i < len(m); i++ {
		ret[i] = '0'
	}
	return string(ret)
}

func defVs(m string) string {
	ret := make([]byte, len(m)-1)
	for i := 1; i < len(m); i++ {
		ret[i-1] = '9'
	}
	return string(ret)
}

// x >= y
func minus(x, y string) string {
	fmt.Println("minus x =", x, "y =", y)
	if x == y {
		fmt.Println("minux ret = 0")
		return "0"
	}
	for i := 0; i < len(x)-len(y); i++ {
		y = "0" + y
	}
	var ( 
		vx, vy byte
		ret []byte
		carry uint8
	)
	for i := len(y)-1; i >= 0; i-- { 
		vx = x[i] - carry
		vy = y[i]
		if vx >= vy {
			ret = append(ret, vx-vy+'0')
			carry = 0
		} else {
			ret = append(ret, vx-vy+10+'0')
			carry = 1
		}
	}

	for i := 0; i <= (len(ret)-1)/2; i++ {
		ret[i], ret[len(ret)-1-i] = ret[len(ret)-1-i], ret[i]
	}
	
	for i := 0; i < len(ret); i++ {
		if ret[i] != '0' {
			ret = ret[i:]
			break
		}
	}

	fmt.Println("minux ret =", string(ret))
	return string(ret)
}
