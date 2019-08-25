//给定两个二进制字符串，返回他们的和（用二进制表示）。
//
//输入为非空字符串且只包含数字 1 和 0。
//
//示例 1:
//
//输入: a = "11", b = "1"
//输出: "100"
//示例 2:
//
//输入: a = "1010", b = "1011"
//输出: "10101"

package add_binary

import (
    "strings"
)

func addBinary(a string, b string) string {
    alen := len(a)
    blen := len(b)
    if alen == 0 {
        return b
    } else if blen == 0 {
        return a
    }
    i := alen-1
    j := blen-1
    carry := 0
    ret := make([]int, 0)
    for i >= 0 && j >= 0 {
        val := int(a[i]- '0' + b[j]-'0')+carry
        carry = val>>1
        ret = append(ret, val%2)
        i--
        j--
    }
    if j >= 0 {
        i = j
        a = b
    }
    for x := i; x >=0; x -- {
        val := int(a[x]-'0') + carry
        carry = val>>1
        ret = append(ret, val%2)
    }
    rlen := len(ret)

    sb := strings.Builder{}
    if carry == 1 {
        sb.WriteString("1")
    }
    for x := rlen-1; x >= 0; x -- {
        sb.WriteByte(byte(ret[x]+'0'))
    }
    return sb.String()
}