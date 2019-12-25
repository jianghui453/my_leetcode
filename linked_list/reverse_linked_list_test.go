package linked_list

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	var head *ListNode
	var hope, ret []int

	toArray := func(node *ListNode) []int {
		r := make([]int, 0)
		for node != nil {
			r = append(r, node.Val)
			node = node.Next
		}
		return r
	}

	head = &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					4,
					&ListNode{
						5,
						nil,
					},
				},
			},
		},
	}
	head = reverseList(head)
	hope = []int{5, 4, 3, 2, 1}
	ret = toArray(head)
	t.Logf("\nhope=%v \n ret=%v", hope, ret)

	head = &ListNode{
		1,
		nil,
	}
	head = reverseList(head)
	hope = []int{1}
	ret = toArray(head)
	t.Logf("\nhope=%v \n ret=%v", hope, ret)

	head = nil
	head = reverseList(head)
	hope = []int{}
	ret = toArray(head)
	t.Logf("\nhope=%v \n ret=%v", hope, ret)
}
