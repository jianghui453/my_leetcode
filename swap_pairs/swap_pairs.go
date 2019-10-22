package swap_pairs

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var last, first, second *ListNode
	first = head
	second = head.Next.Next
	head = head.Next
	head.Next = first
	head.Next.Next = second

	last, first = head.Next, head.Next.Next
	if first == nil {
		return head
	}
	second = first.Next
	for first != nil && second != nil {
		first.Next = second.Next
		last.Next = second
		second.Next = first

		last = first
		first = first.Next
		if first == nil {
			return head
		}
		second = first.Next
	}
	return head
}
