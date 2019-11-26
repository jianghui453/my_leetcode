package tree

import "testing"

func TestSortedArrayToBST(t *testing.T) {
	var nums []int
	var r, h *TreeNode

	var preorder func(node *TreeNode) []int
	preorder = func(node *TreeNode) []int {
		if node == nil {
			return []int{}
		}

		r := make([]int, 0)
		r = append(r, node.Val)
		r = append(r, preorder(node.Left)...)
		r = append(r, preorder(node.Right)...)
		return r
	}

	nums = []int{-10, -3, 0, 5, 9}
	h = &TreeNode{
		0,
		&TreeNode{
			-3,
			&TreeNode{
				-10,
				nil,
				nil,
			},
			nil,
		},
		&TreeNode{
			9,
			&TreeNode{
				5,
				nil,
				nil,
			},
			nil,
		},
	}
	r = sortedArrayToBST(nums)
	t.Logf("\nnums=%v \nh=%v \nr=%v", nums, preorder(h), preorder(r))
}
