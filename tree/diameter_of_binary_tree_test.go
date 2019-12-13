package tree

import (
	"testing"
)

func Test_diameterOfBinaryTree(t *testing.T) {
	var root *TreeNode
	var hope, ret int

	root = &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{4, nil, nil},
			&TreeNode{5, nil, nil},
		},
		&TreeNode{3, nil, nil},
	}
	hope = 3
	ret = diameterOfBinaryTree(root)
	t.Log(ret == hope, hope, ret)

	root = &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{
				4,
				&TreeNode{
					2,
					&TreeNode{4, nil, nil},
					&TreeNode{5, nil, nil},
				},
				nil,
			},
			&TreeNode{
				5,
				&TreeNode{
					2,
					&TreeNode{4, nil, nil},
					&TreeNode{5, nil, nil},
				},
				nil,
			},
		},
		&TreeNode{3, nil, nil},
	}
	hope = 6
	ret = diameterOfBinaryTree(root)
	t.Log(ret == hope, hope, ret)
}