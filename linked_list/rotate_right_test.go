package linked_list

import "testing"

func TestRotateRight(t *testing.T) {
	var head *ListNode
	var k int
	var h, r []int

	var f func(n *ListNode) []int
	f = func(n *ListNode) []int {
		ret := make([]int, 0)
		if n == nil {
			return ret
		}

		ret = append(ret, append([]int{n.Val}, f(n.Next)...)...)
		return ret
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
	k = 2
	h = []int{4,5,1,2,3}
	head = rotateRight(head, k)
	r = f(head)
	t.Logf("\nh=%v \nr=%v", h, r)
}
