package linked_list

import "testing"

func TestDeleteDuplicate2(t *testing.T) {
	var head, ret *ListNode
	var hope, retSlice []int

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
	hope = []int{1, 2, 5}
	retSlice = make([]int, 0)
	ret = deleteDuplicates2(head)
	for ret != nil {
		retSlice = append(retSlice, ret.Val)
		ret = ret.Next
	}
	t.Logf("\n    hope=%v \nretSlice=%v", hope, retSlice)

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
	hope = []int{2, 3}
	retSlice = make([]int, 0)
	ret = deleteDuplicates2(head)
	for ret != nil {
		retSlice = append(retSlice, ret.Val)
		ret = ret.Next
	}
	t.Logf("\n    hope=%v \nretSlice=%v", hope, retSlice)
}
