// 实现一个高效的单向链表逆序输出

package reverse_link

type ListNode struct {
	val  int
	next *ListNode
}

func reverseLink(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	ret := head
	head = head.next
	ret.next = nil
	for head != nil {
		_head := head
		head = head.next
		_head.next = ret
		ret = _head
	}
	return ret
}
