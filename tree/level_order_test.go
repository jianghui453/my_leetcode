package tree

import "testing"

func TestLevelOrderTest(t *testing.T) {
    var root *TreeNode
    var h, r [][]int

    root = &TreeNode{
        3,
        &TreeNode{
            9,
            nil,
            nil,
        },
        &TreeNode{
            20,
            &TreeNode{
                15,
                nil,
                nil,
            },
            &TreeNode{
                7,
                nil,
                nil,
            },
        },
    }
    h = [][]int{{3}, {9,20}, {15,7}}
    r = levelOrder(root)
    t.Logf("h=%v r=%v", h, r)
}

func TestZigzagLevelOrderTest(t *testing.T) {
    var root *TreeNode
    var h, r [][]int

    root = &TreeNode{
        3,
        &TreeNode{
            9,
            nil,
            nil,
        },
        &TreeNode{
            20,
            &TreeNode{
                15,
                nil,
                nil,
            },
            &TreeNode{
                7,
                nil,
                nil,
            },
        },
    }
    h = [][]int{{3}, {20, 9}, {15,7}}
    r = zigzagLevelOrder(root)
    t.Logf("h=%v r=%v", h, r)
}
