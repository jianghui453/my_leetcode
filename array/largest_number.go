// 给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。

// 示例 1:

// 输入: [10,2]
// 输出: 210
// 示例 2:

// 输入: [3,30,34,5,9]
// 输出: 9534330
// 说明: 输出结果可能非常大，所以你需要返回一个字符串而不是整数。

package array

func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}

	var (
		lag []int
	)

	for i := range nums {
		num := nums[i]
		numAry := make([]int, 0)
		for num > 0 {
			numAry = append(numAry, num%10)
			num /= 10
		}

		if len(lag) == 0 {
			lag = numAry
		} else {
			nIdx, lagIdx := len(numAry)-1, len(lag)-1
			for nIdx >= 0 || lagIdx >= 0 {
				if nIdx < 0 {
					if lag[lagIdx] < lag[lagIdx+1] {
						lag = numAry
					}
				} else if lagIdx < 0 {
					if numAry[nIdx] >= numAry[nIdx+1] {
						lag = numAry
					}
				} else if numAry[nIdx] > lag[lagIdx] {
					lag = numAry
				} else {
					nIdx--
					lagIdx--
					continue
				}
				break
			}
		}
	}
}
