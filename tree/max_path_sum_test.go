package tree

import "testing"

func TestMaxPathSum(t *testing.T) {
	var root *TreeNode
	var h, r int

	root = &TreeNode{
		1,
		&TreeNode{
			2,
			nil,
			nil,
		},
		&TreeNode{
			3,
			nil,
			nil,
		},
	}
	h = 6
	r = maxPathSum(root)
	t.Logf("%t h=%d r=%d", h == r, h, r)

	root = &TreeNode{
		-10,
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
	h = 42
	r = maxPathSum(root)
	t.Logf("%t h=%d r=%d", h == r, h, r)

	//[5,4,8,11,null,13,4,7,2,null,null,null,1]
	//                  5
	//               /     \
	//             4         8
	//            /       /     \
	//          11      13         4
	//          /\                   \
	//         7  2                    1
	root = &TreeNode{
		5,
		&TreeNode{
			4,
			&TreeNode{
				11,
				&TreeNode{
					7,
					nil,
					nil,
				},
				&TreeNode{
					2,
					nil,
					nil,
				},
			},
			nil,
		},
		&TreeNode{
			8,
			&TreeNode{
				13,
				nil,
				nil,
			},
			&TreeNode{
				4,
				nil,
				&TreeNode{
					1,
					nil,
					nil,
				},
			},
		},
	}
	h = 48
	r = maxPathSum(root)
	t.Logf("%t h=%d r=%d", h == r, h, r)
}
