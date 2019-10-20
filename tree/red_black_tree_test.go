package tree

import "testing"

func TestRBT_Insert(t *testing.T) {
    var root *RBTNode
    var h, r []int

    rbt := new(RBT)

    root = rbt.Insert(root, 7)
    h = []int{7}
    r = rbt.inOrder(root)
    t.Logf("\ninOrder \nh=%v \nr=%v", h, r)
    h = []int{7}
    r = rbt.levelOrder(root)
    t.Logf("\nlevelOrder \nh=%v \nr=%v", h, r)

    root = rbt.Insert(root, 6)
    h = []int{6, 7}
    r = rbt.inOrder(root)
    t.Logf("\ninOrder \nh=%v \nr=%v", h, r)
    h = []int{7, 6}
    r = rbt.levelOrder(root)
    t.Logf("\nlevelOrder \nh=%v \nr=%v", h, r)

    root = rbt.Insert(root, 5)
    h = []int{5, 6, 7}
    r = rbt.inOrder(root)
    t.Logf("\ninOrder \nh=%v \nr=%v", h, r)
    h = []int{6, 5, 7}
    r = rbt.levelOrder(root)
    t.Logf("\nlevelOrder \nh=%v \nr=%v", h, r)

    root = rbt.Insert(root, 4)
    h = []int{4, 5, 6, 7}
    r = rbt.inOrder(root)
    t.Logf("\ninOrder \nh=%v \nr=%v", h, r)
    h = []int{6, 5, 7, 4}
    r = rbt.levelOrder(root)
    t.Logf("\nlevelOrder \nh=%v \nr=%v", h, r)

    root = rbt.Insert(root, 3)
    h = []int{3, 4, 5, 6, 7}
    r = rbt.inOrder(root)
    t.Logf("\ninOrder \nh=%v \nr=%v", h, r)
    h = []int{6, 4, 7, 3, 5}
    r = rbt.levelOrder(root)
    t.Logf("\nlevelOrder \nh=%v \nr=%v", h, r)

    root = rbt.Insert(root, 2)
    h = []int{2, 3, 4, 5, 6, 7}
    r = rbt.inOrder(root)
    t.Logf("\ninOrder \nh=%v \nr=%v", h, r)
    h = []int{6, 4, 7, 3, 5, 2}
    r = rbt.levelOrder(root)
    t.Logf("\nlevelOrder \nh=%v \nr=%v", h, r)

    root = rbt.Insert(root, 1)
    h = []int{1, 2, 3, 4, 5, 6, 7}
    r = rbt.inOrder(root)
    t.Logf("\ninOrder \nh=%v \nr=%v", h, r)
    h = []int{6, 4, 7, 2, 5, 1, 3}
    r = rbt.levelOrder(root)
    t.Logf("\nlevelOrder \nh=%v \nr=%v", h, r)
}
