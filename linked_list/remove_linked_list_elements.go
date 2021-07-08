// 删除链表中等于给定值 val 的所有节点。

// 示例:

// 输入: 1->2->6->3->4->5->6, val = 6
// 输出: 1->2->3->4->5

package linked_list

func removeElements(head *ListNode, val int) *ListNode {
	var (
		tail, h *ListNode = head, head
	)

	for h != nil {
		if h.Val == val {
			if head == h {
				head, tail, h = head.Next, tail.Next, h.Next
			} else {
				h = h.Next
			}
		} else {
			tail.Next, tail, h = h, h, h.Next
		}
	}

	if tail != nil {
		tail.Next = nil
	}
	return head
}
