package tree

import "testing"

func TestPostOrderTraversal(t *testing.T) {
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
	h = []int{}
	getPostOrderTraversal(root, &h)
	r = postorderTraversal(root)
	t.Logf("\n h=%v \n r=%v", h, r)

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
				nil,
				nil,
			},
		},
		&TreeNode{
			5,
			&TreeNode{
				6,
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
	h = []int{}
	getPostOrderTraversal(root, &h)
	r = postorderTraversal(root)
	t.Logf("\n h=%v \n r=%v", h, r)
}

func getPostOrderTraversal(node *TreeNode, r *[]int) {
	if node.Left != nil {
		getPostOrderTraversal(node.Left, r)
	}
	if node.Right != nil {
		getPostOrderTraversal(node.Right, r)
	}
	*r = append(*r, node.Val)
}
