// 一个强密码应满足以下所有条件：

// 由至少6个，至多20个字符组成。
// 至少包含一个小写字母，一个大写字母，和一个数字。
// 同一字符不能连续出现三次 (比如 "...aaa..." 是不允许的, 但是 "...aa...a..." 是可以的)。
// 编写函数 strongPasswordChecker(s)，s 代表输入字符串，如果 s 已经符合强密码条件，则返回0；否则返回要将 s 修改为满足强密码条件的字符串所需要进行修改的最小步数。

// 插入、删除、替换任一字符都算作一次修改。

package string

func strongPasswordChecker(s string) int {
    var (
		missingUpperCase, missingLowerCase, missingNumber = 1, 1, 1
		modCnt [3]int
		l = len(s)
		modifyCnt, removeCnt, missingTypeCnt int
	)

	for i := 0; i < l; i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			missingLowerCase = 0
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			missingUpperCase = 0
		} else if s[i] >= '0' && s[i] <= '9' {
			missingNumber = 0
		}

		var (
			cnt = 1
			j = i+1
		)
		for ; j < l && s[j] == s[i]; j++ {
			cnt++
		}

		if cnt >= 3 {
			modifyCnt += cnt/3
			modCnt[cnt%3]++
		}
		i = j-1
	}

	missingTypeCnt = missingLowerCase + missingUpperCase + missingNumber

	if l < 6 {
		return max(6-l, missingTypeCnt)
	}

	if l <= 20 {
		return max(missingTypeCnt, modifyCnt)
	}

	removeCnt = l - 20

	// 3n 型子串无法完全变为 3n+2 型，每个需要 1 次删除，只能把 removeCnt 个变为 3n+2 型减少 removeCnt 次后续修改
	if removeCnt < modCnt[0] {
		return max(modifyCnt - removeCnt, missingTypeCnt) + l - 20
	}

	// 3n 型全部变为 3n+2 型 
	removeCnt -= modCnt[0]
	modifyCnt -= modCnt[0]

	// 3n+1 型无法完全变为 3n+2 型，每个需要 2 次删除，减少 removeCnt/2 次后续修改
	if removeCnt < modCnt[1]*2 {
		return max(modifyCnt - removeCnt/2, missingTypeCnt) + l - 20
	}

	removeCnt -= modCnt[1] * 2
	modifyCnt -= modCnt[1]
	
	return max(modifyCnt - removeCnt/3, missingTypeCnt) + l - 20
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
