package linked_list

import "testing"

func TestReverseBetween(t *testing.T) {
	var head, h *ListNode
	var ho, r []int
	var m, n int

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
	ho = []int{1, 4, 3, 2, 5}
	r = []int{}
	h = reverseBetween(head, m, n)
	for h != nil {
		r = append(r, h.Val)
		h = h.Next
	}
	t.Logf("\nho=%v \n r=%v", ho, r)

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
    ho = []int{2, 1, 3, 4, 5}
    r = []int{}
    h = reverseBetween(head, m, n)
    for h != nil {
        r = append(r, h.Val)
        h = h.Next
    }
    t.Logf("\nho=%v \n r=%v", ho, r)

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
    ho = []int{5, 4, 3, 2, 1}
    r = []int{}
    h = reverseBetween(head, m, n)
    for h != nil {
        r = append(r, h.Val)
        h = h.Next
    }
    t.Logf("\nho=%v \n r=%v", ho, r)

    head = &ListNode{
        1,
        nil,
    }
    m = 1
    n = 6
    ho = []int{w1}
    r = []int{}
    h = reverseBetween(head, m, n)
    for h != nil {
        r = append(r, h.Val)
        h = h.Next
    }
    t.Logf("\nho=%v \n r=%v", ho, r)
}
