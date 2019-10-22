package leet_code

import (
	"fmt"
)

// type NodeList struct{
// 	Val int
// 	Next *NodeList
// }

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	var start *ListNode = l1
	if l2 != nil && l2.Val < start.Val {
		start = l2
		l1, l2 = l2, l1
	}
	l1 = l1.Next
	var end *ListNode = start
	for {
		fmt.Printf("start = %v; end = %v; l1 = %v; l2 = %v\n", start, end, l1, l2)
		if l1 == nil {
			end.Next = l2
			break
		} else if l2 == nil {
			end.Next = l1
			break
		}
		if l1.Val <= l2.Val {
			end.Next, l1 = l1, l1.Next
		} else {
			end.Next, l2 = l2, l2.Next
		}
		end = end.Next
	}
	return start
}

func mergeTwoListsV1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head, end *ListNode
	if l1.Val >= l2.Val {
		head = l2
		end = l2
		l2 = l2.Next
	} else {
		head = l1
		end = l1
		l1 = l1.Next
	}
	for {
		if l1 == nil {
			end.Next = l2
			break
		}
		if l2 == nil {
			end.Next = l1
			break
		}
		if l1.Val >= l2.Val {
			end.Next = l2
			l2 = l2.Next
		} else {
			end.Next = l1
			l1 = l1.Next
		}
		end = end.Next
	}
	return head
}
