package tree

import (
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	var root, p, q, hope, ret *TreeNode

	p = &TreeNode{
		5,
		&TreeNode{
			6,
			nil,
			nil,
		},
		&TreeNode{
			2,
			&TreeNode{7, nil, nil},
			&TreeNode{4, nil, nil},
		},
	}
	q = &TreeNode{
		1,
		&TreeNode{0, nil, nil},
		&TreeNode{8, nil, nil},
	}
	root = &TreeNode{
		3,
		p,
		q,
	}
	hope, ret = root, lowestCommonAncestor(root, p, q)
	t.Logf("%t root=%v p=%v q=%v hope=%v ret=%v", ret == hope, root, p, q, hope, ret)

	q = &TreeNode{4, nil, nil}
	p = &TreeNode{
		5,
		&TreeNode{
			6,
			nil,
			nil,
		},
		&TreeNode{
			2,
			&TreeNode{7, nil, nil},
			q,
		},
	}
	root = &TreeNode{
		3,
		p,
		&TreeNode{
			1,
			&TreeNode{0, nil, nil},
			&TreeNode{8, nil, nil},
		},
	}
	hope, ret = p, lowestCommonAncestor(root, p, q)
	t.Logf("%t root=%v p=%v q=%v hope=%v ret=%v", ret == hope, root, p, q, hope, ret)
}
