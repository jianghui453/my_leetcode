//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//示例：
//
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807

//标签：链表

package linked_list

func addTwoNumbersMine(l1, l2 *ListNode) *ListNode {
	carry := 0
	var head, tail *ListNode
	for l1 != nil && l2 != nil {
		val := (l1.Val + l2.Val + carry) % 10
		carry = (l1.Val + l2.Val + carry) / 10
		node := &ListNode{
			val,
			nil,
		}
		if tail == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if l2 != nil {
		l1 = l2
	}
	for l1 != nil {
		val := (l1.Val + carry) % 10
		carry = (l1.Val + carry) / 10
		node := &ListNode{
			val,
			nil,
		}
		if tail == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
		l1 = l1.Next
	}
	if carry != 0 {
		tail.Next = &ListNode{
			carry,
			nil,
		}
	}
	return head
}

//func addTwoNumbersMine(l1, l2 *ListNode) *ListNode {
//	rNL := l1
//Loop:
//	for {
//		l1.Val = l1.Val + l2.Val
//		if l1.Next != nil && l2.Next != nil {
//			if l1.Val >= 10 {
//				l1.Val = l1.Val % 10
//				l1.Next.Val++
//			}
//			l1 = l1.Next
//			l2 = l2.Next
//			continue
//		}
//		if l1.Next == nil {
//			l1.Next = l2.Next
//		}
//		if l1.Val >= 10 {
//			l1.Val = l1.Val % 10
//			if l1.Next == nil {
//				l1.Next = new(ListNode)
//				l1.Next.Val++
//				break
//			}
//			l1.Next.Val++
//			l1 = l1.Next
//		}
//		for {
//			if l1.Val < 10 {
//				break Loop
//			}
//			l1.Val = 0
//			if l1.Next == nil {
//				l1.Next = new(ListNode)
//			}
//			l1 = l1.Next
//			l1.Val++
//		}
//	}
//	return rNL
//}
