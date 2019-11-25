package tree

import "testing"

func TestMaxDepth(t *testing.T) {
	var root *TreeNode
	var h, r int

	root = &TreeNode{
		3,
		&TreeNode{
			9,
			nil,
			nil,
		},
		&TreeNode{
			20,
			&TreeNode{
				15,
				nil,
				nil,
			},
			&TreeNode{
				7,
				nil,
				nil,
			},
		},
	}
	h = 3
	r = maxDepth(root)
	t.Logf("%t h=%d r=%d", h == r, h, r)

	root = &TreeNode{
		3,
		&TreeNode{
			9,
			nil,
			nil,
		},
		&TreeNode{
			20,
			&TreeNode{
				15,
				nil,
				nil,
			},
			&TreeNode{
				21,
				nil,
				&TreeNode{
					22,
					nil,
					nil,
				},
			},
		},
	}
	h = 4
	r = maxDepth(root)
	t.Logf("%t h=%d r=%d", h == r, h, r)

	root = &TreeNode{
		3,
		nil,
		nil,
	}
	h = 1
	r = maxDepth(root)
	t.Logf("%t h=%d r=%d", h == r, h, r)

	root = nil
	h = 0
	r = maxDepth(root)
	t.Logf("%t h=%d r=%d", h == r, h, r)
}
