package tree

import "testing"

func TestMinDepth(t *testing.T) {
	var root *TreeNode
	var h, r int

	root = &TreeNode{
		1,
		nil,
		&TreeNode{
			2,
			nil,
			nil,
		},
	}
	h = 2
	r = minDepth(root)
	t.Logf("%t h=%d r=%d", h == r, h, r)

	root = &TreeNode{
		1,
		nil,
		&TreeNode{
			2,
			nil,
			&TreeNode{
				3,
				nil,
				nil,
			},
		},
	}
	h = 3
	r = minDepth(root)
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
				7,
				nil,
				nil,
			},
		},
	}
	h = 2
	r = minDepth(root)
	t.Logf("%t h=%d r=%d", h == r, h, r)

	root = &TreeNode{
		1,
		nil,
		nil,
	}
	h = 1
	r = minDepth(root)
	t.Logf("%t h=%d r=%d", h == r, h, r)
}
