package linked_list

import "testing"

func TestReverseBetween(t *testing.T) {
	var (
		head      *ListNode
		hope, ret []int
		m, n      int
		f         func(n *ListNode) []int
	)

	f = func(n *ListNode) []int {
		r := make([]int, 0)
		if n == nil {
			return r
		}
		r = append(r, n.Val)
		return append(r, f(n.Next)...)
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
	m = 2
	n = 4
	hope = []int{1, 4, 3, 2, 5}
	head = reverseBetween(head, m, n)
	ret = f(head)
	t.Logf("m=%d n=%d \nhope=%v \n ret=%v", m, n, hope, ret)

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
	m = 1
	n = 2
	hope = []int{2, 1, 3, 4, 5}
	head = reverseBetween(head, m, n)
	ret = f(head)
	t.Logf("m=%d n=%d \nhope=%v \n ret=%v", m, n, hope, ret)

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
	m = 1
	n = 6
	hope = []int{5, 4, 3, 2, 1}
	head = reverseBetween(head, m, n)
	ret = f(head)
	t.Logf("m=%d n=%d \nhope=%v \n ret=%v", m, n, hope, ret)

	head = &ListNode{
		1,
		nil,
	}
	m = 1
	n = 6
	hope = []int{1}
	head = reverseBetween(head, m, n)
	ret = f(head)
	t.Logf("m=%d n=%d \nhope=%v \n ret=%v", m, n, hope, ret)
}
