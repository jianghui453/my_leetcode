package tree

import "fmt"

type BTreeNode struct {
	Keys []int
	C    []*BTreeNode
	T    int // degree of btree
	Leaf bool
	N    int // number of keys
}

func (n *BTreeNode) Init(t int, leaf bool) {
	n.Keys = make([]int, 0)
	n.C = make([]*BTreeNode, 0)
	n.T = t
	n.Leaf = leaf
}

func (n *BTreeNode) Traverse() []int {
	if n != nil {
		fmt.Printf("n%v\n", n)
	}

	ret := make([]int, 0)
	i := 0

	for ; i < n.N; i++ {
		if !n.Leaf {
			ret = append(ret, n.C[i].Traverse()...)
		}

		ret = append(ret, n.Keys[i])
	}

	if !n.Leaf {
		ret = append(ret, n.C[i].Traverse()...)
	}
	return ret
}

func (n *BTreeNode) Search(k int) *BTreeNode {
	i := 0
	for i < n.N && k > n.Keys[i] {
		i++
	}

	if n.Keys[i] == k {
		return n
	}

	if n.Leaf {
		return nil
	}
	return n.C[i].Search(k)
}

func (n *BTreeNode) SplitChild(i int, y *BTreeNode) {
	t := n.T
	z := new(BTreeNode)
	z.Init(y.T, y.Leaf)
	z.N = t - 1

	for j := 0; j < t-1; j++ {
		z.Keys = append(z.Keys, y.Keys[j+t])
	}

	if !y.Leaf {
		for j := 0; j < t; j++ {
			z.C = append(z.C, y.C[j+t])
		}
	}

	n.C = append(n.C, nil)
	for j := n.N; j >= i+1; j-- {
		n.C[j+1] = n.C[j]
	}

	n.C[i+1] = z

	n.Keys = append(n.Keys, 0)
	for j := n.N - 1; j >= i; j-- {
		n.Keys[j+1] = n.Keys[j]
	}

	n.Keys[i] = y.Keys[t-1]
	n.N++

	y.Keys = y.Keys[:t]
	y.N = t - 1
}

func (n *BTreeNode) InsertNotFull(k int) {
	i := n.N - 1

	if n.Leaf {
		n.Keys = append(n.Keys, 0)

		for i >= 0 && n.Keys[i] > k {
			n.Keys[i+1] = n.Keys[i]
			i--
		}

		n.Keys[i+1] = k
		n.N++
	} else {
		for i >= 0 && n.Keys[i] > k {
			i--
		}

		if n.C[i+1].N == 2*n.T-1 {
			n.SplitChild(i+1, n.C[i+1])

			if n.Keys[i+1] < k {
				i++
			}
		}

		n.C[i+1].InsertNotFull(k)
	}
}

type BTree struct {
	Root *BTreeNode
	T    int
}

func (b *BTree) Init(t int) {
	b.T = t
}

func (b *BTree) Insert(k int) {
	if b.Root == nil {
		b.Root = new(BTreeNode)
		b.Root.Init(b.T, true)
		b.Root.Keys = append(b.Root.Keys, k)
		b.Root.N++
	} else {
		if b.Root.N == 2*b.T-1 {
			s := new(BTreeNode)
			s.Init(b.T, false)
			s.C = append(s.C, b.Root)
			s.SplitChild(0, b.Root)

			i := 0

			if s.Keys[0] < k {
				i++
			}

			s.C[i].InsertNotFull(k)

			b.Root = s
		} else {
			b.Root.InsertNotFull(k)
		}
	}
}
