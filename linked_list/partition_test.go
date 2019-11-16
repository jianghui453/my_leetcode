package linked_list

import "testing"

func TestPartition(t *testing.T) {
	var (
		head *ListNode
		x int
		hope, ret []int
		f func(n *ListNode) []int
	)

	f = func(n *ListNode) []int {
		if n == nil {
			return []int{}
		}

		return append([]int{n.Val}, f(n.Next)...)
	}

	head = &ListNode{
	   1,
	   &ListNode{
	       4,
	       &ListNode{
	           3,
	           &ListNode{
	               2,
	               &ListNode{
	                   5,
	                   &ListNode{
	                       2,
	                       nil,
	                   },
	               },
	           },
	       },
	   },
	}
	x = 3
	hope = []int{1, 2, 2, 4, 3, 5}
	head = partition(head, x)
	ret = f(head)
	t.Logf("\nhope=%v \n ret=%v", hope, ret)

	head = &ListNode{
		1,
		nil,
	}
	x = 2
	hope = []int{1}
	head = partition(head, x)
	ret = f(head)
	t.Logf("\nhope=%v \n ret=%v", hope, ret)
}
