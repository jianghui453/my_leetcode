package tree

import "testing"

func TestGenerateTrees(t *testing.T) {
	var n int
	var h, r [][]int
    var ret []*TreeNode

	n = 3
	h = [][]int{
		{1, 0, 3, 2},
		{3, 2, 0, 1},
		{3, 1, 0, 0, 2},
		{2, 1, 3},
		{1, 0, 2, 0, 3},
	}
	ret = generateTrees(n)
	r = make([][]int, 0)
	for i := 0; i < len(ret); i++ {
     r = append(r, preOrder(ret[i]))
	}
	t.Logf("n=%d \nh=%v \nr=%v", n, h, r)

    n = 4
    h = [][]int{
       {1,0,2,0,3,0,4},
       {1,0,2,0,4,3},
       {1,0,3,2,4},
       {1,0,4,2,0,0,3},
       {1,0,4,3,0,2},
       {2,1,3,0,0,0,4},
       {2,1,4,0,0,3},
       {3,1,4,0,2},
       {3,2,4,1},
       {4,1,0,0,2,0,3},
       {4,1,0,0,3,2},
       {4,2,0,1,3},
       {4,3,0,1,0,0,2},
       {4,3,0,2,0,1},
    }
    ret = generateTrees(n)
    r = make([][]int, 0)
    for i := 0; i < len(ret); i++ {
       r = append(r, preOrder(ret[i]))
    }
    t.Logf("n=%d \nh=%v \nr=%v", n, h, r)
}

func preOrder(root *TreeNode) []int {
    r := []int{root.Val}
    s := []*TreeNode{root}
    lenS := 1
    for lenS > 0 {
        stack := make([]*TreeNode, 0)
        stackLen := 0
        for i := 0; i < lenS; i ++ {
            if s[i].Left == nil && s[i].Right == nil {
                continue
            }
            if s[i].Left != nil {
                r = append(r, s[i].Left.Val)
                stack = append(stack, s[i].Left)
                stackLen ++
            } else {
                r = append(r, 0)
            }
            if s[i].Right != nil {
                r = append(r, s[i].Right.Val)
                stack = append(stack, s[i].Right)
                stackLen ++
            } else {
                r = append(r, 0)
            }
        }
        s = stack
        lenS = stackLen
    }
    return r
}
