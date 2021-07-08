package tree

import ()

type BTreeNode struct {
	Keys     []int
	KeysCnt  int
	Children []*BTreeNode
	M        int // order of btree
	IsLeaf   bool
}

func (n *BTreeNode) Constructor(m int, isLeaf bool) {
	n.Keys = make([]int, 0)
	n.Children = make([]*BTreeNode, 0)
	n.M = m
	n.IsLeaf = isLeaf
}

func (n *BTreeNode) Traverse() []int {
	ret := make([]int, 0)
	for i := 0; i < n.KeysCnt; i++ {
		if !n.IsLeaf {
			ret = append(ret, n.Children[i].Traverse()...)
		}
		ret = append(ret, n.Keys[i])
	}

	if !n.IsLeaf {
		ret = append(ret, n.Children[n.KeysCnt].Traverse()...)
	}

	return ret
}

func (n *BTreeNode) Search(k int) *BTreeNode {
	for i := 0; i < n.KeysCnt; i++ {
		if n.Keys[i] == k {
			return n
		} else if n.Keys[i] > k {
			ret := n.Children[i].Search(k)
			if ret != nil {
				return ret
			} else {
				return nil
			}
		}
	}

	ret := n.Children[n.KeysCnt].Search(k)
	return ret
}

func (n *BTreeNode) Insert(k int) {
	if !n.IsLeaf {
		i := 0
		for ; i < n.KeysCnt; i++ {
			if n.Keys[i] == k {
				return
			} else if n.Keys[i] > k {
				break
			}
		}
		n.Children[i].Insert(k)
		if n.Children[i].KeysCnt > n.M {
			n.SplitChild(i)
		}
		return
	}

	for i := 0; i < n.KeysCnt; i++ {
		if n.Keys[i] == k {
			return
		} else if n.Keys[i] > k {
			n.Keys = append(n.Keys[:i], append([]int{k}, n.Keys[i:]...)...)
			break
		}
	}
	if len(n.Keys) == n.KeysCnt {
		n.Keys = append(n.Keys, k)
	}
	n.KeysCnt++
}

func (n *BTreeNode) SplitChild(i int) {
	child1 := new(BTreeNode)
	child1.Constructor(n.M, n.Children[i].IsLeaf)
	child2 := new(BTreeNode)
	child2.Constructor(n.M, n.Children[i].IsLeaf)

	child := n.Children[i]
	mid := child.KeysCnt / 2
	for j := 0; j < mid; j++ {
		child1.Keys = append(child1.Keys, child.Keys[j])
		child1.KeysCnt++
		if !child.IsLeaf {
			child1.Children = append(child1.Children, child.Children[j])
		}
	}
	if !child.IsLeaf {
		child1.Children = append(child1.Children, child.Children[mid])
	}

	for j := mid + 1; j < child.KeysCnt; j++ {
		child2.Keys = append(child2.Keys, child.Keys[j])
		child2.KeysCnt++
		if !child.IsLeaf {
			child2.Children = append(child2.Children, child.Children[j])
		}
	}
	if !child.IsLeaf {
		child2.Children = append(child2.Children, child.Children[child.KeysCnt])
	}

	n.Keys = append(n.Keys, child.Keys[mid])
	n.KeysCnt++
	if i == len(n.Children) {
		n.Children = append(n.Children[:i], []*BTreeNode{child1, child2}...)
	} else {
		n.Children = append(n.Children[:i], append([]*BTreeNode{child1, child2}, n.Children[i+1:]...)...)
	}
}

type BTree struct {
	Root *BTreeNode
	M    int
}

func (b *BTree) Constructor(m int) {
	b.M = m
}

func (b *BTree) Insert(k int) {
	if b.Root == nil {
		b.Root = new(BTreeNode)
		b.Root.Constructor(b.M, true)
		b.Root.Keys = append(b.Root.Keys, k)
		b.Root.KeysCnt++
	} else {
		b.Root.Insert(k)
		if b.Root.KeysCnt > b.M {
			root := new(BTreeNode)
			root.Constructor(b.M, false)
			root.Children = []*BTreeNode{b.Root}
			root.SplitChild(0)
			b.Root = root
		}
	}
}
