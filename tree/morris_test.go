package tree

import "testing"

func TestInOrder(t *testing.T) {
	var root *TreeNode
	var h, r []int

	root = &TreeNode{
		6,
		&TreeNode{
			2,
			&TreeNode{
				1,
				nil,
				nil,
			},
			&TreeNode{
				4,
				&TreeNode{
					3,
					nil,
					nil,
				},
				&TreeNode{
					5,
					nil,
					nil,
				},
			},
		},
		&TreeNode{
			7,
			nil,
			&TreeNode{
				9,
				&TreeNode{
					8,
					nil,
					nil,
				},
				nil,
			},
		},
	}
	h = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	r = inOrder(root)
	t.Logf("\nh=%v \nr=%v", h, r)
}

func TestPreOrder(t *testing.T) {
	var root *TreeNode
	var h, r []int

	root = &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{
				3,
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
					6,
					nil,
					nil,
				},
			},
		},
		&TreeNode{
			7,
			nil,
			&TreeNode{
				8,
				&TreeNode{
					9,
					nil,
					nil,
				},
				nil,
			},
		},
	}
	h = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	r = preOrder(root)
	t.Logf("\nh=%v \nr=%v", h, r)
}
