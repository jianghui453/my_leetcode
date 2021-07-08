package linked_list

import "testing"

func TestReorderList(t *testing.T) {
	var head *ListNode
	var h, r []int

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
	r = []int{}
	h = []int{1, 5, 2, 4, 3}
	reorderList(head)
	for head != nil {
		r = append(r, head.Val)
		head = head.Next
	}
	t.Logf("\n h=%v \n r=%v", h, r)

	head = &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					4,
					nil,
				},
			},
		},
	}
	h = []int{1, 4, 2, 3}
	r = []int{}
	reorderList(head)
	for head != nil {
		r = append(r, head.Val)
		head = head.Next
	}
	t.Logf("\n h=%v \n r=%v", h, r)
}
