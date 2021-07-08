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
	S []*TreeNode
	Node *TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		Node: root,
	}
}


func (this *BSTIterator) Next() int {
	for n := this.Node; n != nil; n = n.Left {
		this.S = append(this.S, n)
	}
	this.Node, this.S = this.S[len(this.S)-1], this.S[:len(this.S)-1]
	ret := this.Node.Val
	this.Node = this.Node.Right
	return ret
}


func (this *BSTIterator) HasNext() bool {
	return this.Node != nil || len(this.S) > 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */