package tree

import "testing"

func TestIsValidBST(t *testing.T) {
	var root *TreeNode
	var h, r bool

	root = &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}
	h = true
	r = isValidBST(root)
	t.Logf("%t h=%t r=%t", h == r, h, r)

	root = &TreeNode{5, &TreeNode{1, nil, nil}, &TreeNode{6,
		&TreeNode{2, nil, nil}, &TreeNode{7, nil, nil}}}
	h = false
	r = isValidBST(root)
	t.Logf("%t h=%t r=%t", h == r, h, r)

	root = &TreeNode{5, &TreeNode{1, nil, nil}, &TreeNode{7,
		&TreeNode{6, nil, nil}, &TreeNode{8, nil, nil}}}
	h = true
	r = isValidBST(root)
	t.Logf("%t h=%t r=%t", h == r, h, r)

	root = &TreeNode{1, &TreeNode{1, nil, nil}, nil}
	h = false
	r = isValidBST(root)
	t.Logf("%t h=%t r=%t", h == r, h, r)

	root = &TreeNode{1, &TreeNode{0, nil, nil}, nil}
	h = true
	r = isValidBST(root)
	t.Logf("%t h=%t r=%t", h == r, h, r)
}
