package reverse_k_group

// import (
// 	"fmt"
// )

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
    if k == 1 {
		return head
	}
	listLen := 0
	current, tail := head, head
	for current != nil {
		listLen ++
		current = current.Next
	}
	for listLen >= k {
		var subHead, subTail *ListNode
		if tail == head {
			subHead, subTail = tail, tail
		} else {
			subHead, subTail = tail.Next, tail.Next
		}
		for i := 0; i < k - 1; i ++ {
			temp := subTail.Next
			subTail.Next = temp.Next
			temp.Next = subHead
			subHead = temp
		}
		if tail == head {
			head = subHead
		} else {
			tail.Next = subHead
		}
		tail = subTail
		listLen -= k
	}
	return head
}