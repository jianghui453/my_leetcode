// 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

package leet_code

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbersMine(l1, l2 *ListNode) *ListNode {
	rNL := l1
Loop:
	for {
		l1.Val = l1.Val + l2.Val
		if l1.Next != nil && l2.Next != nil {
			if l1.Val >= 10 {
				l1.Val = l1.Val % 10
				l1.Next.Val++
			}
			l1 = l1.Next
			l2 = l2.Next
			continue
		}
		if l1.Next == nil {
			l1.Next = l2.Next
		}
		if l1.Val >= 10 {
			l1.Val = l1.Val % 10
			if l1.Next == nil {
				l1.Next = new(ListNode)
				l1.Next.Val++
				break
			}
			l1.Next.Val++
			l1 = l1.Next
		}
		for {
			if l1.Val < 10 {
				break Loop
			}
			l1.Val = 0
			if l1.Next == nil {
				l1.Next = new(ListNode)
			}
			l1 = l1.Next
			l1.Val++
		}
	}
	return rNL
}
