package tree

import "testing"

func TestInorderIPAddress(t *testing.T) {
	var root *TreeNode
	var h, r []int

	root = &TreeNode{
		1,
		nil,
		&TreeNode{
			2,
			&TreeNode{
				3,
				nil,
				nil,
			},
			nil,
		},
	}
	h = []int{1, 3, 2}
	r = inorderTraversal(root)
	t.Logf("root=%v \nh=%v \nr=%v", root, h, r)

    root = &TreeNode{
        1,
        nil,
        nil,
    }
    h = []int{1}
    r = inorderTraversal(root)
    t.Logf("root=%v \nh=%v \nr=%v", root, h, r)

    root = &TreeNode{
        1,
        &TreeNode{
            2,
            &TreeNode{
                3,
                nil,
                nil,
            },
            nil,
        },
        nil,
    }
    h = []int{3, 2, 1}
    r = inorderTraversal(root)
    t.Logf("root=%v \nh=%v \nr=%v", root, h, r)

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
    h = []int{1, 2, 3}
    r = inorderTraversal(root)
    t.Logf("root=%v \nh=%v \nr=%v", root, h, r)
}
