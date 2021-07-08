// 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
// k 是一个正整数，它的值小于或等于链表的长度。
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
// 示例 :
// 给定这个链表：1->2->3->4->5
// 当 k = 2 时，应当返回: 2->1->4->3->5
// 当 k = 3 时，应当返回: 3->2->1->4->5
// 说明 :
// 你的算法只能使用常数的额外空间。
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

package linked_list

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 || head == nil || head.Next == nil {
		return head
	}

	rcd := make([][]*ListNode, k)
	for i := 0; i < k; i++ {
		rcd[i] = make([]*ListNode, 0)
	}

	h := head
	for h != nil {
		for i := 0; i < k; i++ {
			if h == nil {
				break
			}
			rcd[i] = append(rcd[i], h)
			h = h.Next
		}
	}

	head = nil
	var tail *ListNode
	l := 0
	for {
		if l >= len(rcd[k-1]) {
			for i := 0; i < k; i++ {
				if l >= len(rcd[i]) {
					break
				}

				if head == nil {
					head, tail = rcd[i][l], rcd[i][l]
				} else {
					tail.Next = rcd[i][l]
					tail = tail.Next
				}
			}
			break
		}

		for i := k - 1; i >= 0; i-- {
			if head == nil {
				head, tail = rcd[i][l], rcd[i][l]
			} else {
				tail.Next = rcd[i][l]
				tail = tail.Next
			}
		}
		l++
	}
	tail.Next = nil

	return head
}

// func reverseKGroup(head *ListNode, k int) *ListNode {
// 	if k == 1 {
// 		return head
// 	}
// 	listLen := 0
// 	current, tail := head, head
// 	for current != nil {
// 		listLen++
// 		current = current.Next
// 	}
// 	for listLen >= k {
// 		var subHead, subTail *ListNode
// 		if tail == head {
// 			subHead, subTail = tail, tail
// 		} else {
// 			subHead, subTail = tail.Next, tail.Next
// 		}
// 		for i := 0; i < k-1; i++ {
// 			temp := subTail.Next
// 			subTail.Next = temp.Next
// 			temp.Next = subHead
// 			subHead = temp
// 		}
// 		if tail == head {
// 			head = subHead
// 		} else {
// 			tail.Next = subHead
// 		}
// 		tail = subTail
// 		listLen -= k
// 	}
// 	return head
// }
