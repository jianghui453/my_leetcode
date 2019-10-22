package tree

func inOrder(root *TreeNode) []int {
	cur := root
	var prev *TreeNode
	r := make([]int, 0)
	for cur != nil {
		if cur.Left == nil {
			r = append(r, cur.Val)
			cur = cur.Right
		} else {
			prev = cur.Left
			for prev.Right != nil && prev.Right != cur {
				prev = prev.Right
			}
			if prev.Right == nil {
				prev.Right = cur
				cur = cur.Left
			} else {
				prev.Right = nil
				r = append(r, cur.Val)
				cur = cur.Right
			}
		}
	}
	return r
}

func preOrder(root *TreeNode) []int {
	cur := root
	var prev *TreeNode
	r := make([]int, 0)
	for cur != nil {
		if cur.Left == nil {
			r = append(r, cur.Val)
			cur = cur.Right
		} else {
			prev = cur.Left
			for prev.Right != nil && prev.Right != cur {
				prev = prev.Right
			}
			if prev.Right == nil {
				prev.Right = cur
				r = append(r, cur.Val)
				cur = cur.Left
			} else {
				cur = cur.Right
				prev.Right = nil
			}
		}
	}
	return r
}
