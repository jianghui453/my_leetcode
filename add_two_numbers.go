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
