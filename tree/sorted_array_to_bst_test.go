package tree

import "testing"

func TestSortedArrayToBST(t *testing.T) {
	var nums []int
	var r, h *TreeNode
	//var f func(n *TreeNode) []int
	//f = func (n *TreeNode) []int {
	//    var r []int
	//    var _f func(_n *TreeNode)
	//    _f = func (_n *TreeNode) {
	//        if _n.Left != nil {
	//            _f(_n.Left)
	//        } else {
	//            r = append(r, 0)
	//        }
	//        r = append(r, _n.Val)
	//        if _n.Right != nil {
	//            _f(_n.Right)
	//        } else {
	//            r = append(r, 0)
	//        }
	//    }
	//    _f(n)
	//    return r
	//}

	nums = []int{-10, -3, 0, 5, 9}
	h = &TreeNode{
		0,
		&TreeNode{
			-3,
			&TreeNode{
				-10,
				nil,
				nil,
			},
			nil,
		},
		&TreeNode{
			9,
			&TreeNode{
				5,
				nil,
				nil,
			},
			nil,
		},
	}
	r = sortedArrayToBST(nums)
	t.Logf("\nnums=%v \nh=%v \nr=%v", nums, printTree(h), printTree(r))
}

func printTree(n *TreeNode) []int {
	var r []int
	var f func(_n *TreeNode, idx int)
	f = func(_n *TreeNode, idx int) {
		if _n.Left != nil {
			if idx*2+1 > len(r)-1 {
				r = append(r, make([]int, idx*2+1-len(r)+1)...)
			}
			r[idx*2+1] = _n.Left.Val
			f(_n.Left, idx*2+1)
		}
		if _n.Right != nil {
			if idx*2+2 > len(r)-1 {
				r = append(r, make([]int, idx*2+2-len(r)+1)...)
			}
			r[idx*2+2] = _n.Right.Val
			f(_n.Right, idx*2+2)
		}
	}
	r = append(r, n.Val)
	f(n, 0)
	return r
}
