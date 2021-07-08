import pprint

def test():
    root = Node(1, [
        Node(3, [
            Node(5, None),
            Node(6, None)
        ]),
        Node(2, None), 
        Node(4, None)
    ])
    c = Codec()
    s = c.serialize(root)
    pprint.pprint(s)

test()
