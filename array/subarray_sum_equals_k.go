// 给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。

// 示例 1 :

// 输入:nums = [1,1,1], k = 2
// 输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
// 说明 :

// 数组的长度为 [1, 20,000]。
// 数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。

package array

// "fmt"

// 这种方法背后的想法如下：如果累积总和（由 sum[i]sum[i] 表示加到 i^{th}ith 的和）最多两个指数是相同的，那么这些元素之间的元素总和为零。
// 进一步扩展相同的想法，如果累计总和，在索引 ii 和 jj 处相差 kk，即 sum[i] - sum[j] = ksum[i]−sum[j]=k，则位于索引 ii 和 jj 之间的元素之和是 kk。
// 基于这些想法，可以使用了一个哈希表 mapmap，它用于存储所有可能的索引的累积总和以及相同累加和发生的次数。我们以以下形式存储数据：(sum_isum i，sum_isum i 的出现次数)。
// 我们遍历数组numsnums并继续寻找累积总和。每当我们遇到一个新的和时，我们在hashmap中创建一个与该总和相对应的新条目。如果再次出现相同的和，我们增加与mapmap中的和相对应的计数。
// 此外，对于遇到的每个总和，我们还确定已经发生 sum-ksum−k 总和的次数，因为它将确定具有总和 kk 的子阵列发生到当前索引的次数。我们将 countcount 增加相同的量。
// 在完成便利数组后，countcount 记录了所需结果
func subarraySum(nums []int, k int) int {
	res, sum := 0, 0
	m := make(map[int]int, len(nums))
	m[0] = 1
	for i := range nums {
		sum += nums[i]
		res += m[sum-k]
		m[sum]++
	}
	return res
}

// func subarraySum(nums []int, k int) int {
// 	var (
// 		l = len(nums)
// 		ret int
// 		prefixSum = make([]int, l)
// 	)

// 	if l == 0 {
// 		return 0
// 	}

// 	prefixSum[0] = nums[0]
// 	for i := 1; i < l; i++ {
// 		prefixSum[i] = prefixSum[i-1] + nums[i]
// 	}

// 	fmt.Println(prefixSum)

// 	for i := 0; i < l; i++ {
// 		for j := i; j < l; j++ {
// 			if prefixSum[j]-prefixSum[i]+nums[i] == k {
// 				fmt.Println(i, j)
// 				ret++
// 			}
// 		}
// 	}

// 	return ret
// }
