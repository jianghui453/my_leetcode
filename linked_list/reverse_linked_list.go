// 反转一个单链表。

// 示例:

// 输入: 1->2->3->4->5->NULL
// 输出: 5->4->3->2->1->NULL
// 进阶:
// 你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

package linked_list

// func reverseList(head *ListNode) *ListNode {
// 	var f func(node *ListNode) (*ListNode, *ListNode)
// 	f = func(node *ListNode) (*ListNode, *ListNode) {
// 		if node == nil || node.Next == nil {
// 			return node, node
// 		}

// 		h, t := f(node.Next)
// 		t.Next, node.Next = node, nil
// 		return h, node
// 	}

// 	head, _ = f(head)
// 	return head
// }

func reverseList(head *ListNode) *ListNode {
	s := make([]*ListNode, 0)
	for head != nil {
		s = append(s, head)
		head = head.Next
	}
	var t *ListNode
	for i := range s {
		if head == nil {
			head, t = s[i], s[i]
		} else {
			s[i].Next, head = head, s[i]
		}
	}
	if t != nil {
		t.Next = nil
	}
	return head
}
