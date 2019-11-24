package tree

import (
	"testing"
)

func TestRecoverTree(t *testing.T) {
	var root, h *TreeNode
	//var h, r []int

	//root = &TreeNode{
	//	1,
	//	&TreeNode{
	//		3,
	//		nil,
	//		&TreeNode{
	//			2,
	//			nil,
	//			nil,
	//		},
	//	},
	//	nil,
	//}
	//h = &TreeNode{
	//	3,
	//	&TreeNode{
	//		1,
	//		nil,
	//		&TreeNode{
	//			2,
	//			nil,
	//			nil,
	//		},
	//	},
	//	nil,
	//}
	//t.Logf("\nr=%v", inOrderT(root))
	//recoverTree(root)
	//t.Logf("\nh=%v \nr=%v", inOrderT(h), inOrderT(root))

	//[10,5,15,0,8,13,20,2,-5,6,9,12,14,18,25]

	root = &TreeNode{
		10,
		&TreeNode{
			5,
			&TreeNode{
				0,
				&TreeNode{
					2,
					nil,
					nil,
				},
				&TreeNode{
					-5,
					nil,
					nil,
				},
			},
			&TreeNode{
				8,
				&TreeNode{
					6,
					nil,
					nil,
				},
				&TreeNode{
					9,
					nil,
					nil,
				},
			},
		},
		&TreeNode{
			15,
			&TreeNode{
				13,
				&TreeNode{
					12,
					nil,
					nil,
				},
				&TreeNode{
					14,
					nil,
					nil,
				},
			},
			&TreeNode{
				20,
				&TreeNode{
					18,
					nil,
					nil,
				},
				&TreeNode{
					25,
					nil,
					nil,
				},
			},
		},
	}
	h = &TreeNode{
		10,
		&TreeNode{
			5,
			&TreeNode{
				0,
				&TreeNode{
					-5,
					nil,
					nil,
				},
				&TreeNode{
					2,
					nil,
					nil,
				},
			},
			&TreeNode{
				8,
				&TreeNode{
					6,
					nil,
					nil,
				},
				&TreeNode{
					9,
					nil,
					nil,
				},
			},
		},
		&TreeNode{
			15,
			&TreeNode{
				13,
				&TreeNode{
					12,
					nil,
					nil,
				},
				&TreeNode{
					14,
					nil,
					nil,
				},
			},
			&TreeNode{
				20,
				&TreeNode{
					18,
					nil,
					nil,
				},
				&TreeNode{
					25,
					nil,
					nil,
				},
			},
		},
	}
	t.Logf("\nr=%v", inOrderT(root))
	recoverTree(root)
	t.Logf("\nh=%v \nr=%v", inOrderT(h), inOrderT(root))
}

func inOrderT(root *TreeNode) []int {
	r := make([]int, 0)
	s := make([]*TreeNode, 0)
	for root != nil {
		if root.Left == nil {
			r = append(r, root.Val)
			if root.Right == nil && len(s) > 0 {
				root = s[len(s)-1]
				s = s[:len(s)-1]
				r = append(r, root.Val)
			}
			root = root.Right
		} else {
			s = append(s, root)
			root = root.Left
		}
	}
	return r
}
