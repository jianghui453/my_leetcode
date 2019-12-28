// 返回 A 的最短的非空连续子数组的长度，该子数组的和至少为 K 。

// 如果没有和至少为 K 的非空子数组，返回 -1 。

//  

// 示例 1：

// 输入：A = [1], K = 1
// 输出：1
// 示例 2：

// 输入：A = [1,2], K = 4
// 输出：-1
// 示例 3：

// 输入：A = [2,-1,2], K = 3
// 输出：3
//  

// 提示：

// 1 <= A.length <= 50000
// -10 ^ 5 <= A[i] <= 10 ^ 5
// 1 <= K <= 10 ^ 9

package sliding_window

import (
	// "fmt"
)

// class Solution(object):
//     def shortestSubarray(self, A, K):
//         N = len(A)
//         P = [0]
//         for x in A:
//             P.append(P[-1] + x)

//         #Want smallest y-x with Py - Px >= K
//         ans = N+1 # N+1 is impossible
//         monoq = collections.deque() #opt(y) candidates, represented as indices of P
//         for y, Py in enumerate(P):
//             #Want opt(y) = largest x with Px <= Py - K
//             while monoq and Py <= P[monoq[-1]]:
//                 monoq.pop()

//             while monoq and Py - P[monoq[0]] >= K:
//                 ans = min(ans, y - monoq.popleft())

//             monoq.append(y)

//         return ans if ans < N+1 else -1

func shortestSubarray(A []int, K int) int {
	var (
		N int = len(A)
		P []int = make([]int, 0)
		ans int
		monoq []int = make([])
	)

	for i := range A {
		if i == 0 {
			P = append(P, 0)
		} else {
			P = append(P, P[i-1]+A[i-1])
		}
	}

	ans = N+1
	monoq = 

	if ans == l+1 {
		return -1
	}
	return ans
}

func setAns(ans, length int) int {
	if ans == -1 || ans > length {
		ans = length
	}
}