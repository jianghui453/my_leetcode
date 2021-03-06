package tree

import "testing"

func TestPathSum(t *testing.T) {
	var root *TreeNode
	var sum int
	var h, r bool

	root = &TreeNode{
		5,
		&TreeNode{
			4,
			&TreeNode{
				11,
				&TreeNode{
					7,
					nil,
					nil,
				},
				&TreeNode{
					2,
					nil,
					nil,
				},
			},
			nil,
		},
		&TreeNode{
			8,
			&TreeNode{
				13,
				nil,
				nil,
			},
			&TreeNode{
				4,
				&TreeNode{
					5,
					nil,
					nil,
				},
				&TreeNode{
					1,
					nil,
					nil,
				},
			},
		},
	}
	sum = 22
	h = true
	r = pathSum(root, sum)
	t.Logf("%t h=%t r=%t", r == h, h, r)
}
