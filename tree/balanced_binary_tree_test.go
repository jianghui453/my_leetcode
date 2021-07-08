package tree

import (
	"testing"
)

func TestIsBalanced(t *testing.T) {
	var root *TreeNode
	var hope, ret bool

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
	hope = true
	ret = isBalanced(root)
	t.Logf("%t hope=%t ret=%t", ret == hope, hope, ret)

	root = &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{
				3,
				&TreeNode{
					4,
					nil,
					nil,
				},
				&TreeNode{
					4,
					nil,
					nil,
				},
			},
			&TreeNode{
				3,
				nil,
				nil,
			},
		},
		&TreeNode{
			2,
			nil,
			nil,
		},
	}
	hope = false
	ret = isBalanced(root)
	t.Logf("%t hope=%t ret=%t", ret == hope, hope, ret)
}
