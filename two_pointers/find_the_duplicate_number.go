// 给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。

// 示例 1:

// 输入: [1,3,4,2,2]
// 输出: 2
// 示例 2:

// 输入: [3,1,3,4,2]
// 输出: 3
// 说明：

// 不能更改原数组（假设数组是只读的）。
// 只能使用额外的 O(1) 的空间。
// 时间复杂度小于 O(n2) 。
// 数组中只有一个重复的数字，但它可能不止重复出现一次。

package two_pointers

import (
	// "fmt"
)

// 1、因为值在 1 到 n 之间，所以 nums[0] 的值不会是 0，所以 nums[nums[0]] 的值不会和键相同，以此类推，不存在自循环。
// 2、因为其中有重复的值，所以当第二个重复的值出现的时候就指向了之前出现过的数，此时出现了环。
// 3、通过快慢指针法找到环的入口，入口键就是重复的值。
func findDuplicate(nums []int) int {
	slow := nums[nums[0]]
	fast := nums[slow]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	ptr := nums[0]
	for slow != ptr {
		slow = nums[slow]
		ptr = nums[ptr]
	}

	return slow
}
