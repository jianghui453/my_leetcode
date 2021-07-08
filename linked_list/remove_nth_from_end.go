// 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
// 示例：
// 给定一个链表: 1->2->3->4->5, 和 n = 2.
// 当删除了倒数第二个节点后，链表变为 1->2->3->5.
// 说明：
// 给定的 n 保证是有效的。
// 进阶：
// 你能尝试使用一趟扫描实现吗？
package linked_list

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n == 0 {
		return head
	}
	his := make([]*ListNode, 0)
	h := head
	for h != nil {
		his = append(his, h)
		h = h.Next
	}
	if n == len(his) {
		if len(his) > 1 {
			return his[1]
		}
		return nil
	}
	if n == 1 {
		his[len(his)-n-1].Next = nil
	} else {
		his[len(his)-n-1].Next = his[len(his)-n+1]
	}
	return head
}

// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	var current *ListNode = head
// 	var record = make([]*ListNode, n+1, n+1)
// 	for {
// 		record = append([]*ListNode{current}, record[:n+1]...)
// 		if current.Next == nil {
// 			if record[n] == nil {
// 				head = record[n-1].Next
// 			} else {
// 				record[n].Next = record[n-1].Next
// 			}
// 			break
// 		}
// 		current = current.Next
// 	}
// 	return head
// }
