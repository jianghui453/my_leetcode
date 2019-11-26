package tree

import "testing"

func TestBuildTree(t *testing.T) {
	var inorder, postorder []int
	var node *TreeNode

	var getInorder func(n *TreeNode) []int
	getInorder = func(n *TreeNode) []int {
		r := make([]int, 0)
		if n == nil {
			return r
		}

		r = append(r, getInorder(n.Left)...)
		r = append(r, n.Val)
		r = append(r, getInorder(n.Right)...)
		return r
	}

	var getPreorder func(n *TreeNode) []int
	getPreorder = func(n *TreeNode) []int {
		r := make([]int, 0)
		if n == nil {
			return r
		}

		r = append(r, n.Val)
		r = append(r, getPreorder(n.Left)...)
		r = append(r, getPreorder(n.Right)...)
		return r
	}

	var getPostorder func(n *TreeNode) []int
	getPostorder = func(n *TreeNode) []int {
		r := make([]int, 0)
		if n == nil {
			return r
		}

		r = append(r, getPostorder(n.Left)...)
		r = append(r, getPostorder(n.Right)...)
		r = append(r, n.Val)
		return r
	}

	inorder = []int{9, 3, 15, 20, 7}
	postorder = []int{9, 15, 7, 20, 3}
	node = buildTree(inorder, postorder)
	t.Logf("\ninorder=%v \n    ret=%v", inorder, getInorder(node))
	t.Logf("\npostorder=%v \n    ret=%v", postorder, getPostorder(node))

	inorder = []int{9}
	postorder = []int{9}
	node = buildTree(inorder, postorder)
	t.Logf("\ninorder=%v \n    ret=%v", inorder, getInorder(node))
	t.Logf("\npostorder=%v \n    ret=%v", postorder, getPostorder(node))

	inorder = []int{9, 15}
	postorder = []int{9, 15}
	node = buildTree(inorder, postorder)
	t.Logf("\ninorder=%v \n    ret=%v", inorder, getInorder(node))
	t.Logf("\npostorder=%v \n    ret=%v", postorder, getPostorder(node))
}

// func TestBuildTree(t *testing.T) {
// 	var inorder, postorder []int
// 	var r []int
// 	var _r *TreeNode

// 	postorder = []int{9, 15, 7, 20, 3}
// 	inorder = []int{9, 3, 15, 20, 7}
// 	_r = buildTree(inorder, postorder)
// 	r = getInorder(_r)
// 	t.Logf("\nh=%v \nr=%v", inorder, r)
// }
