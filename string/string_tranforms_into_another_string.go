// 给出两个长度相同的字符串，分别是 str1 和 str2。请你帮忙判断字符串 str1 能不能在 零次 或 多次 转化后变成字符串 str2。

// 每一次转化时，将会一次性将 str1 中出现的 所有 相同字母变成其他 任何 小写英文字母（见示例）。

// 只有在字符串 str1 能够通过上述方式顺利转化为字符串 str2 时才能返回 True，否则返回 False。​​

//  

// 示例 1：

// 输入：str1 = "aabcc", str2 = "ccdee"
// 输出：true
// 解释：将 'c' 变成 'e'，然后把 'b' 变成 'd'，接着再把 'a' 变成 'c'。注意，转化的顺序也很重要。
// 示例 2：

// 输入：str1 = "leetcode", str2 = "codeleet"
// 输出：false
// 解释：我们没有办法能够把 str1 转化为 str2。
//  

// 提示：

// 1 <= str1.length == str2.length <= 10^4
// str1 和 str2 中都只会出现 小写英文字母

package string

import (
	"fmt"
)

func canConvert(str1 string, str2 string) bool {
	if str1 == str2 {
		return true
	}

	var (
		l1, l2 = len(str1), len(str2)
		hash1, hash2 [26]int
		rule [26]int
		m1 = make(map[int]int)
		m2 = make(map[int]int)
	)
	
	if l1 != l2 {
		return false
	}

	for i := 0; i < l1; i++ {
		k1, k2 := int(str1[i] - 'a'), int(str2[i] - 'a')

		if rule[k1] == 0 {
			rule[k1] = k2 + 1
		} else if rule[k1] != k2 + 1 {
			fmt.Println(1)
			return false
		}

		hash1[k1]++
		hash2[k2]++
	}

	// fmt.Println(hash1, hash2)

	for i := 0; i < 26; i++ {
		m1[hash1[i]]++
		m2[hash2[i]]++
	}

	// fmt.Println(m1, m2)

	for _, v := range m2 {
		if v == 26 {
			return false
		}
	}

	return true
}
