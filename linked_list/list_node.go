package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func ToArray(node *ListNode) []int {
	var (
		ret []int
	)

	for node != nil {
		ret = append(ret, node.Val)
		node = node.Next
	}

	return ret
}
