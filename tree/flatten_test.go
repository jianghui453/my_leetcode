package tree

import "testing"

func TestFlatter(t *testing.T) {
    var root *TreeNode
    var h []int

    root = &TreeNode{
        1,
        &TreeNode{
            2,
            &TreeNode{
                3,
                nil,
                nil,
            },
            &TreeNode{
                4,
                nil,
                nil,
            },
        },
        &TreeNode{
            5,
            nil,
            &TreeNode{
                6,
                nil,
                nil,
            },
        },
    }
    h = []int{1, 2, 3, 4, 5, 6}
    flatten(root)
    t.Logf("\nh=%v \nr=%v", h, readLink(root))
}

func readLink(node *TreeNode) []int {
    r := make([]int, 0)
    for node != nil {
        r = append(r, node.Val)
        node = node.Right
    }
    return r
}
