// 请判断一个链表是否为回文链表。

// 示例 1:

// 输入: 1->2
// 输出: false
// 示例 2:

// 输入: 1->2->2->1
// 输出: true
// 进阶：
// 你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

package linked_list

import (
// "fmt"
)

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	var h *ListNode
	slow, fast := head, head.Next
	for {
		if fast.Next == nil {
			h, slow.Next, slow = slow, h, slow.Next
			break
		}
		if fast.Next.Next == nil {
			h, slow.Next, slow = slow, h, slow.Next.Next
			break
		}

		slow, fast, h, slow.Next = slow.Next, fast.Next.Next, slow, h
	}

	for h != nil {
		if h.Val != slow.Val {
			return false
		}

		h, slow = h.Next, slow.Next
	}

	return true
}
