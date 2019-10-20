package tree

import "testing"

func TestAVL_Insert(t *testing.T) {
    var root *AVLNode
    var h, r []int

    avl := new(AVL)

    root = avl.Insert(root, 10)
    h = []int{10}
    r = avl.preOrder(root)
    t.Logf("\nh=%v \nr=%v", h, r)

    root = avl.Insert(root, 20)
    h = []int{10, 20}
    r = avl.preOrder(root)
    t.Logf("\nh=%v \nr=%v", h, r)

    root = avl.Insert(root, 30)
    h = []int{20, 10, 30}
    r = avl.preOrder(root)
    t.Logf("\nh=%v \nr=%v", h, r)

    root = avl.Insert(root, 40)
    h = []int{20, 10, 30, 40}
    r = avl.preOrder(root)
    t.Logf("\nh=%v \nr=%v", h, r)

    root = avl.Insert(root, 50)
    h = []int{20, 10, 40, 30, 50}
    r = avl.preOrder(root)
    t.Logf("\nh=%v \nr=%v", h, r)

    root = avl.Insert(root, 25)
    h = []int{30, 20, 10, 25, 40, 50}
    r = avl.preOrder(root)
    t.Logf("\nh=%v \nr=%v", h, r)

    root = avl.Delete(root, 20)
    h = []int{30, 10, 25, 40, 50}
    r = avl.preOrder(root)
    t.Logf("\nh=%v \nr=%v", h, r)


    //root = avl.Insert(root, 9)
    //root = avl.Insert(root, 5)
    //root = avl.Insert(root, 10)
    //root = avl.Insert(root, 0)
    //root = avl.Insert(root, 6)
    //root = avl.Insert(root, 11)
    //root = avl.Insert(root, -1)
    //root = avl.Insert(root, 1)
    //root = avl.Insert(root, 2)
    //h = []int{9, 1, 0, -1, 5, 2, 6, 10, 11}
    //r = avl.preOrder(root)
    //t.Logf("\nh=%v \nr=%v", h, r)
    //
    //root = avl.Delete(root, 10)
    //h = []int{1, 0, -1, 9, 5, 2, 6, 11}
    //r = avl.preOrder(root)
    //t.Logf("\nh=%v \nr=%v", h, r)
}
