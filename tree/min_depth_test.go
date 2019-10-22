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
}
