package tree

import "testing"

func TestBTree_Insert(t *testing.T) {
	var btree *BTree
	var h, r []int

	btree = new(BTree)
	btree.Constructor(3)

	btree.Insert(10)
	h = []int{10}
	r = btree.Root.Traverse()
	t.Logf("\nh=%v \nr=%v", h, r)

	btree.Insert(20)
	h = []int{10, 20}
	r = btree.Root.Traverse()
	t.Logf("\nh=%v \nr=%v", h, r)

	btree.Insert(5)
	h = []int{5, 10, 20}
	r = btree.Root.Traverse()
	t.Logf("\nh=%v \nr=%v", h, r)

	btree.Insert(6)
	h = []int{5, 6, 10, 20}
	r = btree.Root.Traverse()
	t.Logf("\nh=%v \nr=%v", h, r)

	btree.Insert(12)
	h = []int{5, 6, 10, 12, 20}
	r = btree.Root.Traverse()
	t.Logf("\nh=%v \nr=%v", h, r)

	btree.Insert(30)
	h = []int{5, 6, 10, 12, 20, 30}
	r = btree.Root.Traverse()
	t.Logf("\nh=%v \nr=%v", h, r)

	btree.Insert(7)
	h = []int{5, 6, 7, 10, 12, 20, 30}
	r = btree.Root.Traverse()
	t.Logf("\nh=%v \nr=%v", h, r)

	btree.Insert(17)
	h = []int{5, 6, 7, 10, 12, 17, 20, 30}
	r = btree.Root.Traverse()
	t.Logf("\nh=%v \nr=%v", h, r)
}
