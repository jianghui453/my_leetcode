// union-find algorithm with rank and compress path optimize


package union_find

type UnionFind struct {
    V int
    parent []int
    rank []int
}

func (u *UnionFind) Init () {
    for i := 0; i < u.V; i ++ {
        u.parent[i] = i
        u.rank[i] = 0
    }
}

func (u *UnionFind) Find (x int) int {
    if u.parent[x] != x {
        u.parent[x] = u.Find(u.parent[x])
    }
    return u.parent[x]
}

func (u *UnionFind) Union (x, y int) {
    xSet := u.Find(x)
    ySet := u.Find(y)
    if xSet == ySet {
        return
    }
    if u.rank[xSet] > u.rank[ySet] {
        u.parent[ySet] = xSet
    } else {
        u.parent[xSet] = ySet
        if u.rank[xSet] == u.rank[ySet] {
            u.rank[ySet] ++
        }
    }
}

func (u *UnionFind) Same (x, y int) bool {
    return u.Find(x) == u.Find(y)
}