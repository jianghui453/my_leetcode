//运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
//
//获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
//写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。
//
//进阶:
//
//你是否可以在 O(1) 时间复杂度内完成这两种操作？
//
//示例:
//
//LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
//
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // 返回  1
//cache.put(3, 3);    // 该操作会使得密钥 2 作废
//cache.get(2);       // 返回 -1 (未找到)
//cache.put(4, 4);    // 该操作会使得密钥 1 作废
//cache.get(1);       // 返回 -1 (未找到)
//cache.get(3);       // 返回  3
//cache.get(4);       // 返回  4

package design

type LRUCache struct {
    cap int
    size int
    cache map[int]*DLinkList
    head *DLinkList
    tail *DLinkList
}

type DLinkList struct {
    Key int
    Val int
    Prev *DLinkList
    Next *DLinkList
}

func Constructor(capacity int) LRUCache {
    l := LRUCache{
        capacity,
        0,
        make(map[int]*DLinkList),
        new(DLinkList),
        new(DLinkList),
    }
    l.head.Prev = l.tail
    l.head.Next = l.tail
    l.tail.Next = l.head
    l.tail.Prev = l.head
    return l
}

func (l *LRUCache) addNode(node *DLinkList) {
    next := l.head.Next
    l.head.Next = node
    node.Prev = l.head
    node.Next = next
    next.Prev = node
}

func (l *LRUCache) rmNode(node *DLinkList) {
    prev := node.Prev
    next := node.Next
    prev.Next = next
    next.Prev = prev
}

func (l *LRUCache) popNode() *DLinkList {
    node := l.tail.Prev
    prev := l.tail.Prev.Prev
    l.tail.Prev = prev
    prev.Next = l.tail
    return node
}

func (this *LRUCache) Get(key int) int {
    if _, ok := this.cache[key]; !ok {
        return -1
    }
    node := this.cache[key]
    this.rmNode(node)
    this.addNode(node)
    return node.Val
}

func (this *LRUCache) Put(key int, value int)  {
    if _, ok := this.cache[key]; ok {
        node := this.cache[key]
        this.rmNode(node)
        this.addNode(node)
        node.Val = value
        return
    }
    node := &DLinkList{
        key,
        value,
        nil,
        nil,
    }
    this.cache[key] = node
    this.size ++
    this.addNode(node)
    if this.size > this.cap {
        node = this.popNode()
        this.size --
        delete(this.cache, node.Key)
    }
}
