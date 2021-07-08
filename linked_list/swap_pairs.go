// 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
// 示例:
// 给定 1->2->3->4, 你应该返回 2->1->4->3.

package linked_list

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd := make([]*ListNode, 0)
	even := make([]*ListNode, 0)

	i := 0
	for head != nil {
		if i%2 == 0 {
			even = append(even, head)
		} else {
			odd = append(odd, head)
		}
		head = head.Next
		i++
	}

	head = odd[0]
	head.Next = even[0]
	tail := head.Next
	for i = 1; i < len(odd); i++ {
		tail.Next = odd[i]
		tail = tail.Next
		if i < len(even) {
			tail.Next = even[i]
			tail = tail.Next
		}
	}
	if i < len(even) {
		tail.Next = even[i]
		tail = tail.Next
	}
	tail.Next = nil

	return head
}

// func swapPairs(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	if head.Next == nil {
// 		return head
// 	}
// 	var last, first, second *ListNode
// 	first = head
// 	second = head.Next.Next
// 	head = head.Next
// 	head.Next = first
// 	head.Next.Next = second

// 	last, first = head.Next, head.Next.Next
// 	if first == nil {
// 		return head
// 	}
// 	second = first.Next
// 	for first != nil && second != nil {
// 		first.Next = second.Next
// 		last.Next = second
// 		second.Next = first

// 		last = first
// 		first = first.Next
// 		if first == nil {
// 			return head
// 		}
// 		second = first.Next
// 	}
// 	return head
// }
