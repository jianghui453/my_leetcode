package tree

import "testing"

func TestSumNumbers(t *testing.T) {
    var root *TreeNode
    var h, r int

    root = &TreeNode{
        1,
        &TreeNode{
            2,
            nil,
            nil,
        },
        &TreeNode{
            3,
            nil,
            nil,
        },
    }
    h = 25
    r = sumNumbers(root)
    t.Logf("%t h=%d r=%d", h==r, h, r)

    root = &TreeNode{
        4,
        &TreeNode{
            9,
            &TreeNode{
                5,
                nil,
                nil,
            },
            &TreeNode{
                1,
                nil,
                nil,
            },
        },
        &TreeNode{
            0,
            nil,
            nil,
        },
    }
    h = 1026
    r = sumNumbers(root)
    t.Logf("%t h=%d r=%d", h==r, h, r)
}
