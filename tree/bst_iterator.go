package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	Root *TreeNode
	Cur *TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		Root root
		Cur root
	}
}


func (this *BSTIterator) Next() int {
	node := this.Root
	for node != nil {
		if node.Val == Cur.Val {
			if node.Right != nil {
				for n := node.Right {
					if n.Left != nil {
						n = n.Left
					} else {
						this.Cur = n
						return n.Val
					}
				}
			} else {
				for i := len(his) - 1; i >= 0; i-- {

				}
			}			
		}
	}
}


func (this *BSTIterator) HasNext() bool {
	node := this.Root
	for node != nil {
		if Cur.Val <= node.Val {
			return node.Right != nil
		}
		if Cur.Val > node.Val {
			node = node.Right
		}
	}
	return false
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */