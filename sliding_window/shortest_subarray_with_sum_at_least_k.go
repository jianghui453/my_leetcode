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

// 解法：
// 如果 x1 < x2 且 P[x2] <= P[x1]，那么 opt(y) 的值不可能为 x1，这是因为 x2 比 x1 大，并且如果 x1 满足了 P[x1] <= P[y] - K，那么 P[x2] <= P[x1] <= P[y] - K，即 x2 同样满足 P[x2] <= P[y] - K。
// 如果 opt(y1) 的值为 x，那么我们以后就不用再考虑 x 了。这是因为如果有 y2 > y1 且 opt(y2) 的值也为 x，但此时 y2 - x 显然大于 y1 - x，不会作为所有 y - opt(y) 中的最小值。

package sliding_window

import (
// "fmt"
)

// class Solution {
//     public int shortestSubarray(int[] A, int K) {
//         int N = A.length;
//         long[] P = new long[N+1];
//         for (int i = 0; i < N; ++i)
//             P[i+1] = P[i] + (long) A[i];

//         // Want smallest y-x with P[y] - P[x] >= K
//         int ans = N+1; // N+1 is impossible
//         Deque<Integer> monoq = new LinkedList(); //opt(y) candidates, as indices of P

//         for (int y = 0; y < P.length; ++y) {
//             // Want opt(y) = largest x with P[x] <= P[y] - K;
//             while (!monoq.isEmpty() && P[y] <= P[monoq.getLast()])
//                 monoq.removeLast();
//             while (!monoq.isEmpty() && P[y] >= P[monoq.getFirst()] + K)
//                 ans = Math.min(ans, y - monoq.removeFirst());

//             monoq.addLast(y);
//         }

//         return ans < N+1 ? ans : -1;
//     }
// }

func shortestSubarray(A []int, K int) int {
	var (
		l         int   = len(A)
		prefixSum []int = make([]int, l+1)
		ret       int   = l + 1
		dqueue          = make([]int, 0)
	)

	for i := 1; i <= l; i++ {
		prefixSum[i] = prefixSum[i-1] + A[i-1]
	}

	for i := 0; i <= l; i++ {
		// 如果该位置的前缀和小于前面某个位置的前缀和则说明中间这一段的和为负数，后续没必要再对这几个位置进行比较
		for len(dqueue) > 0 && prefixSum[i] <= prefixSum[dqueue[len(dqueue)-1]] {
			dqueue = dqueue[:len(dqueue)-1]
		}

		// 如果左右指针的前缀和差大于k则记录并把左指针往右移一位
		for len(dqueue) > 0 && prefixSum[i]-prefixSum[dqueue[0]] >= K {
			ret = min(ret, i-dqueue[0])
			if len(dqueue) > 1 {
				dqueue = dqueue[1:]
			} else {
				dqueue = dqueue[:0]
			}
		}

		dqueue = append(dqueue, i)
	}

	if ret == l+1 {
		return -1
	}
	return ret
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
