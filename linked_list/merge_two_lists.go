// 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
// 示例：
// 输入：1->2->4, 1->3->4
// 输出：1->1->2->3->4->4

package linked_list

// import (
// 	"fmt"
// )

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, h *ListNode
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			l1, l2 = l2, l1
		}
		if h == nil {
			h = l1
			head = h
		} else {
			h.Next = l1
			h = h.Next
		}
		l1 = l1.Next
	}
	if l1 != nil {
		l1, l2 = l2, l1
	}
	for l2 != nil {
		if h == nil {
			h = l2
			head = h
		} else {
			h.Next = l2
			h = h.Next
		}
		l2 = l2.Next
	}
	if h != nil {
		h.Next = nil
	}
	return head
}

// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
// 	if l1 == nil {
// 		return l2
// 	}
// 	var start *ListNode = l1
// 	if l2 != nil && l2.Val < start.Val {
// 		start = l2
// 		l1, l2 = l2, l1
// 	}
// 	l1 = l1.Next
// 	var end *ListNode = start
// 	for {
// 		fmt.Printf("start = %v; end = %v; l1 = %v; l2 = %v\n", start, end, l1, l2)
// 		if l1 == nil {
// 			end.Next = l2
// 			break
// 		} else if l2 == nil {
// 			end.Next = l1
// 			break
// 		}
// 		if l1.Val <= l2.Val {
// 			end.Next, l1 = l1, l1.Next
// 		} else {
// 			end.Next, l2 = l2, l2.Next
// 		}
// 		end = end.Next
// 	}
// 	return start
// }

// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
// 	if l1 == nil {
// 		return l2
// 	}
// 	if l2 == nil {
// 		return l1
// 	}
// 	var head, end *ListNode
// 	if l1.Val >= l2.Val {
// 		head = l2
// 		end = l2
// 		l2 = l2.Next
// 	} else {
// 		head = l1
// 		end = l1
// 		l1 = l1.Next
// 	}
// 	for {
// 		if l1 == nil {
// 			end.Next = l2
// 			break
// 		}
// 		if l2 == nil {
// 			end.Next = l1
// 			break
// 		}
// 		if l1.Val >= l2.Val {
// 			end.Next = l2
// 			l2 = l2.Next
// 		} else {
// 			end.Next = l1
// 			l1 = l1.Next
// 		}
// 		end = end.Next
// 	}
// 	return head
// }
