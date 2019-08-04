package kth_smallest

import "testing"

func TestKthSmallest(t *testing.T) {
	var root *BinTree
	var k, ret int

	root = &BinTree{
		5,
		&BinTree{
			3,
			&BinTree{
				2,
				&BinTree{
					1,
					nil,
					nil,
				},
				nil,
			},
			&BinTree{
				4,
				nil,
				nil,
			},
		},
	 	&BinTree{
			6,
			nil,
			nil,
		},
	}
	k = 2
	ret = kthSmallest(root, k)
	t.Logf("k=%d ret=%d", k, ret)
}