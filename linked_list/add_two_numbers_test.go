package linked_list

import (
	"testing"
)

func Test_addTwoNumbersMine(t *testing.T) {
	var l1, l2 *ListNode
	var h, r []int

	l1 = &ListNode{
		2,
		&ListNode{
			4,
			&ListNode{
				3,
				nil,
			},
		},
	}
	l2 = &ListNode{
		5,
		&ListNode{
			6,
			&ListNode{
				4,
				nil,
			},
		},
	}
	l1 = addTwoNumbersMine(l1, l2)
	h = []int{7, 0, 8}
	r = listToArray(l1)
	t.Logf("\nh=%v \nr=%v\n", h, r)

	l1 = &ListNode{
		2,
		&ListNode{
			4,
			&ListNode{
				3,
				nil,
			},
		},
	}
	l2 = nil
	l1 = addTwoNumbersMine(l1, l2)
	h = []int{2, 4, 3}
	r = listToArray(l1)
	t.Logf("\nh=%v \nr=%v\n", h, r)

	l1 = &ListNode{
		2,
		&ListNode{
			4,
			&ListNode{
				3,
				nil,
			},
		},
	}
	l2 = &ListNode{
		5,
		&ListNode{
			6,
			nil,
		},
	}
	l1 = addTwoNumbersMine(l1, l2)
	h = []int{7, 0, 4}
	r = listToArray(l1)
	t.Logf("\nh=%v \nr=%v\n", h, r)

	l1 = nil
	l2 = nil
	l1 = addTwoNumbersMine(l1, l2)
	h = []int{}
	r = listToArray(l1)
	t.Logf("\nh=%v \nr=%v\n", h, r)

	l1 = &ListNode{
		9,
		&ListNode{
			9,
			&ListNode{
				9,
				nil,
			},
		},
	}
	l2 = &ListNode{
		9,
		&ListNode{
			9,
			&ListNode{
				9,
				nil,
			},
		},
	}
	l1 = addTwoNumbersMine(l1, l2)
	h = []int{8, 9, 9, 1}
	r = listToArray(l1)
	t.Logf("\nh=%v \nr=%v\n", h, r)

	l1 = &ListNode{
		9,
		&ListNode{
			9,
			nil,
		},
	}
	l2 = &ListNode{
		1,
		nil,
	}
	l1 = addTwoNumbersMine(l1, l2)
	h = []int{0, 0, 1}
	r = listToArray(l1)
	t.Logf("\nh=%v \nr=%v\n", h, r)
}

func listToArray(head *ListNode) []int {
	r := make([]int, 0)
	for head != nil {
		r = append(r, head.Val)
		head = head.Next
	}
	return r
}
