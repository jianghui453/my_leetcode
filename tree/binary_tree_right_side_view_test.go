package tree

import (
	"testing"
)

func Test_rightSideView(t *testing.T) {
	var root *TreeNode
	var hope, ret []int

	root = &TreeNode{
		1,
		&TreeNode{
			2,
			nil,
			&TreeNode{
				5,
				nil,
				nil,
			},
		},
		&TreeNode{
			3,
			nil,
			&TreeNode{
				4,
				nil,
				nil,
			},
		},
	}
	hope, ret = []int{1, 3, 4}, rightSideView(root)
	t.Logf("\nhope=%v \n ret=%v", hope, ret)

	root = &TreeNode{
		1,
		&TreeNode{
			2,
			nil,
			&TreeNode{
				5,
				&TreeNode{
					4,
					nil,
					nil,
				},
				nil,
			},
		},
		&TreeNode{
			3,
			nil,
			&TreeNode{
				4,
				nil,
				nil,
			},
		},
	}
	hope, ret = []int{1, 3, 4, 4}, rightSideView(root)
	t.Logf("\nhope=%v \n ret=%v", hope, ret)
}