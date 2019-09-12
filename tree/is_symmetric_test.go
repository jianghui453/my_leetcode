package tree

import "testing"

func TestIsSymmetric(t *testing.T) {
    var root *TreeNode
    var h, r bool

    //root = &TreeNode{
    //    1,
    //    &TreeNode{
    //        2,
    //        nil,
    //        &TreeNode{
    //            3,
    //            nil,
    //            nil,
    //        },
    //    },
    //    &TreeNode{
    //        2,
    //        nil,
    //        &TreeNode{
    //            3,
    //            nil,
    //            nil,
    //        },
    //    },
    //}
    //h = false
    //r = isSymmetric(root)
    //t.Logf("%t h=%t r=%t", h==r, h, r)
    //
    //root = &TreeNode{
    //   1,
    //   &TreeNode{
    //       2,
    //       &TreeNode{
    //           3,
    //           nil,
    //           nil,
    //       },
    //       &TreeNode{
    //           4,
    //           nil,
    //           nil,
    //       },
    //   },
    //   &TreeNode{
    //       2,
    //       &TreeNode{
    //           4,
    //           nil,
    //           nil,
    //       },
    //       &TreeNode{
    //           3,
    //           nil,
    //           nil,
    //       },
    //   },
    //}
    //h = true
    //r = isSymmetric(root)
    //t.Logf("%t h=%t r=%t", h==r, h, r)

    root = &TreeNode{
        2,
        &TreeNode{
            3,
            &TreeNode{
                4,
                nil,
                nil,
            },
            &TreeNode{
                5,
                nil,
                nil,
            },
        },
        &TreeNode{
            3,
            &TreeNode{
                5,
                nil,
                nil,
            },
            nil,
        },
    }
    h = false
    r = isSymmetric(root)
    t.Logf("%t h=%t r=%t", h==r, h, r)
}
