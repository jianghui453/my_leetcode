package linked_list

import "testing"

func TestDeleteDuplicate(t *testing.T) {
	var head *ListNode
	var hope, ret []int
	var f func(n *ListNode) []int

	f = func(n *ListNode) []int {
		r := make([]int, 0)
		if n == nil {
			return r
		}

		r = append(r, n.Val)
		r = append(r, f(n.Next)...)
		return r
	}

	head = &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					3,
					&ListNode{
						4,
						&ListNode{
							4,
							&ListNode{
								5,
								nil,
							},
						},
					},
				},
			},
		},
	}
	hope = []int{1, 2, 3, 4, 5}
	head = deleteDuplicates(head)
	ret = f(head)
	t.Logf("\nhope=%v \n ret=%v", hope, ret)

	head = &ListNode{
		1,
		&ListNode{
			1,
			&ListNode{
				1,
				&ListNode{
					2,
					&ListNode{
						3,
						nil,
					},
				},
			},
		},
	}
	hope = []int{1, 2, 3}
	head = deleteDuplicates(head)
	ret = f(head)
	t.Logf("\nhope=%v \n ret=%v", hope, ret)

	head = &ListNode{
		1,
		&ListNode{
			1,
			nil,
		},
	}
	hope = []int{1}
	head = deleteDuplicates(head)
	ret = f(head)
	t.Logf("\nhope=%v \n ret=%v", hope, ret)
}
