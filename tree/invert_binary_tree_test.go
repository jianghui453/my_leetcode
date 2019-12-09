package tree

import (
	"testing"
)

func Test_invertTree(t *testing.T) {
	var root *TreeNode
	var hope, ret []int

	root = &TreeNode{
		4,
		&TreeNode{
			2,
			&TreeNode{1, nil, nil},
			&TreeNode{3, nil, nil},
		},
		&TreeNode{
			7,
			&TreeNode{6, nil, nil},
			&TreeNode{9, nil, nil},
		},
	}
	hope = []int{9, 7, 6, 4, 3, 2, 1}
	root = invertTree(root)
	ret = inorderTraversal(root)
	t.Logf("\nhope=%v \n ret=%v", hope, ret)
}