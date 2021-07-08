//反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
//
//说明:
//1 ≤ m ≤ n ≤ 链表长度。
//
//示例:
//
//输入: 1->2->3->4->5->NULL, m = 2, n = 4
//输出: 1->4->3->2->5->NULL

package linked_list

import (
	"math"
)

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	s := make([]*ListNode, 0)
	for head != nil {
		s = append(s, head)
		head = head.Next
	}

	var tail *ListNode
	if m > 1 {
		head, tail = s[0], s[0]
		for i := 1; i < m-1; i++ {
			tail.Next = s[i]
			tail = tail.Next
		}
	}

	for i := int(math.Min(float64(n-1), float64(len(s)-1))); i >= m-1; i-- {
		if head == nil {
			head, tail = s[i], s[i]
		} else {
			tail.Next = s[i]
			tail = tail.Next
		}
	}

	if n >= len(s) {
		tail.Next = nil
	} else {
		tail.Next = s[n]
	}

	return head
}

// func reverseBetween(head *ListNode, m int, n int) *ListNode {
// 	if head == nil {
// 		return head
// 	}
// 	if n <= m {
// 		return head
// 	}
// 	var h1, h2, h3, t1, t2 *ListNode
// 	if m > 1 {
// 		h1 = head
// 		t1 = head
// 	}
// 	for i := 1; i < m-1 && t1 != nil; i++ {
// 		t1 = t1.Next
// 	}
// 	if t1 != nil {
// 		h2 = t1.Next
// 		t2 = h2
// 	} else {
// 		h2 = head
// 		t2 = head
// 	}
// 	h3 = h2.Next
// 	for i := 1; i < n-m+1 && h3 != nil; i++ {
// 		h3 = h3.Next
// 		c := t2.Next
// 		t2.Next = h3
// 		c.Next = h2
// 		h2 = c
// 		if t1 != nil {
// 			t1.Next = h2
// 		}
// 	}
// 	if h1 == nil {
// 		return h2
// 	}
// 	return h1
// }
