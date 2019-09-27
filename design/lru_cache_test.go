package design

import "testing"

func TestLRUCache(t *testing.T) {
    var cache LRUCache
    var output, hope []int

    cache = Constructor(2)
    output = make([]int, 0)
    hope = []int{1, -1, -1, 3, 4}
    cache.Put(1, 1)
    cache.Put(2, 2)
    output = append(output, cache.Get(1))       // 返回  1
    cache.Put(3, 3)    // 该操作会使得密钥 2 作废
    output = append(output, cache.Get(2))       // 返回 -1 (未找到)
    cache.Put(4, 4)    // 该操作会使得密钥 1 作废
    output = append(output, cache.Get(1))       // 返回 -1 (未找到)
    output = append(output, cache.Get(3))       // 返回  3
    output = append(output, cache.Get(4))       // 返回  4
    t.Logf("\noutput=%v \n  hope=%v\n", output, hope)

    //["LRUCache","put","put","get","put","put","get"]
    //[[2],[2,1],[2,2],[2],[1,1],[4,1],[2]]
    cache = Constructor(2)
    output = make([]int, 0)
    hope = []int{2, -1}
    cache.Put(2, 1)
    cache.Put(2, 2)
    output = append(output, cache.Get(2))
    cache.Put(1, 1)
    cache.Put(4, 1)
    output = append(output, cache.Get(2))
    t.Logf("\noutput=%v \n  hope=%v\n", output, hope)
}
