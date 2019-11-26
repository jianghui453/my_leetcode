package tree

import "testing"

func TestLevelOrderTest(t *testing.T) {
	var root *TreeNode
	var h, r [][]int

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
	h = [][]int{{15, 7}, {9, 20}, {3}}
	r = levelOrderBottom(root)
	t.Logf("\nh=%v \nr=%v", h, r)
}
