import pprint

def test():
    pprint.pprint("test")
    root = TreeNode{
        val: 10,
        left: TreeNode{
            val: 5,
            left: TreeNode{
                val: 2,
                left: TreeNode{
                    val: 1,
                    left: None,
                    right: None
                },
                right: TreeNode{
                    val: 3,
                    left: None,
                    right: None
                }
            },
            right: TreeNode{
                val: 15,
                left: TreeNode{
                    val: 12,
                    left: TreeNode{
                        val: 11,
                        left: None,
                        right: None
                    },
                    right: TreeNode{
                        val: 13,
                        left: None,
                        right: None
                    }
                },
                right: None
            }
        },
        right: None
    }
    c = Codec()
    str = c.serialize(root)
    pprint.pprint(str)

test()