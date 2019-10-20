package tree

type BTreeNode struct {
	Keys []int
	C []*BTreeNode
	T int // degree of btree
	Leaf bool
	N int // number of keys
}

func (bn *BTreeNode) new(t int, leaf bool) *BTreeNode {
	n := new(BTreeNode)
	n.T = t
	n.Leaf = leaf
	return n
}

func (bn *BTreeNode) traverse() []int {
	ret := make([]int, 0)
	for i := 0; i < bn.N; i ++ {
		if !bn.Leaf {
			ret = append(ret, bn.C[i].traverse()...)
			ret = append(ret, bn.Keys[i])
		}
	}

}

type BTree struct {
	Root *BTreeNode
	T int
}

func newBTree(t int) *BTree {
	b := new(BTree)
	b.T = t
	return b
}

func (b *BTree) Insert(k int) {
	if b.Root == nil {
		b.Root = new(BTreeNode)
		b.Root.Leaf = true
		b.Root.T = b.T
	} else {

	}
}
