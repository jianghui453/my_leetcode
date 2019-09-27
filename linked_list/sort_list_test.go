package linked_list

import (
    "testing"
)

func TestSortList(t *testing.T) {
    var head *ListNode
    var h, r []int

    head = &ListNode{
        4,
        &ListNode{
            2,
            &ListNode{
                1,
                &ListNode{
                    3,
                    nil,
                },
            },
        },
    }
    h = []int{1, 2, 3, 4}
    head = sortList(head)
    r = ToArray(head)
    t.Logf("\n h=%v \n r=%v", h, r)
}

func ToArray(head *ListNode) []int {
    r := make([]int, 0)
    for head != nil {
        r = append(r, head.Val)
        head = head.Next
    }
    return r
}
