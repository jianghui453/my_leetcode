package leet_code

// type ListNode {
// 	val int
// 	next *ListNode
// }

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var current *ListNode = head
	var record = make([]*ListNode, n+1, n+1)
	for {
		record = append([]*ListNode{current}, record[:n+1]...)
		if current.Next == nil {
			if record[n] == nil {
				head = record[n-1].Next
			} else {
				record[n].Next = record[n-1].Next
			}
			break
		}
		current = current.Next
	}
	return head
}
