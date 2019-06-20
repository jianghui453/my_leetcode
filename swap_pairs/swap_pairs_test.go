package swap_pairs

import(
	"testing"
)

func TestSwapPairs(t *testing.T) {
	head := &ListNode{
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
	ret := swapPairs(head)
	for ret != nil {
		t.Log(ret.Val)
		ret = ret.Next
	}
}