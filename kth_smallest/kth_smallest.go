package kth_smallest

import "fmt"

type BinTree struct {
	val   int
	left  *BinTree
	right *BinTree
}

func kthSmallest(root *BinTree, k int) int {
	if k < 1 {
		return -1
	}
	ret := -1
	search(root, &k, &ret)
	return ret
}

func search(root *BinTree, k, ret *int) {
	fmt.Printf("root.val=%d k=%d\n", root.val, *k)
	if root.left != nil {
		search(root.left, k, ret)
	}
	*k --
	if *k == 0 {
		*ret = root.val
		return
	}
	if root.right != nil {
		search(root.right, k, ret)
	}
}
